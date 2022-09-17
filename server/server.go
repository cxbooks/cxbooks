package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/tools"
	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/gin-gonic/gin"
)

var (
	srv      *Service  //全局单例
	initOnce sync.Once //

)

type Store = model.Store

type Service struct {
	srv  *http.Server //http server
	sSrv *http.Server //https server
	cfg  *Config      //

	router *gin.Engine //http 路由表

	scanner *tools.ScannerManager

	orm *Store

	ctx context.Context
}

func NewService() *Service {

	initOnce.Do(func() {
		config := InitConfig()
		srv = NewServiceWithConfig(config)
	})
	//返回全局的service，主程序启动应该只一定一个service
	return srv

}

// NewServiceWithConfig with config
func NewServiceWithConfig(cf *Config) *Service {

	//opendb
	if cf == nil || (cf.DBOpt == nil) {
		panic(`数据库配置异常，HOST空异常`)
	}

	_, err := os.Stat(cf.DataPath)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.D(`缓存目录`, cf.DataPath, `无法访问或者不存在 `, err)
		panic(`数据库配置异常，HOST空异常`)
	}

	web := &Service{
		cfg: cf,
		// mg: NewManagerWithConfig(cf),
	}

	return web

}

// StartContext 启动
func (m *Service) StartContext(ctx context.Context) (err error) {
	m.ctx = ctx

	go func() {
		//连接数据库
		m.orm, err = model.WaitDB(ctx, m.cfg.DBOpt)
		if err != nil {
			zlog.E(`初始化数据库异常，退出`)
			os.Exit(1)
		}

		//初始化数据库
		if m.cfg.InitFlag {
			if err := model.AutoMigrate(m.orm); err != nil {
				zlog.E(`初始化数据库异常，退出`)
				os.Exit(1)
			}
			os.Exit(0)
		}

		//开启golang 调试
		if m.cfg.PPROF != "" {
			zlog.I(`开启PPROF：`, m.cfg.PPROF)
			go func() {
				zlog.I(http.ListenAndServe(m.cfg.PPROF, nil))
			}()
		}

		//开启http/https服务
		//TODO
		m.router = initGinRoute(m.cfg.LogLevel)

		m.scanner, _ = tools.NewScannerManager(m.ctx, m.cfg.DataPath, m.orm)

		m.listen()
	}()

	return nil

}

// GracefulStop 退出，每个模块实现stop
func (m *Service) GracefulStop() {
	if m.srv != nil {
		zlog.D(`退出HTTP服务...`)
		m.srv.Shutdown(m.ctx)
	}

	if m.sSrv != nil {
		zlog.D(`退出HTTPS服务...`)
		m.srv.Shutdown(m.ctx)
	}

	if m.orm != nil {
		zlog.I(`关闭数据库连接...`)
		m.orm.Close()
	}

	if m.scanner != nil {
		m.scanner.Stop()
	}

}

func (m *Service) Store() *Store {
	return m.orm
}

func (m *Service) listen() {

	//只开启http
	if m.cfg.TLSAddr == "" {
		// m.srv = &http.Server{Addr: m.cfg.APIAddr, Handler: m.router.SetTrustedCIDRs()}
		m.srv = &http.Server{Addr: m.cfg.APIAddr, Handler: m.router}
		zlog.I(`开启HTTP server：`, m.cfg.APIAddr)
		go func() { m.srv.ListenAndServe() }()

		return
	}

	//走到这里，说明 https 一定是开启的 ，这是判断一下是否要开启http
	if m.cfg.APIAddr == "" { //如果APIAddr为空则只开启https

		m.srv = &http.Server{
			Addr:    m.cfg.TLSAddr,
			Handler: m.router,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				},
			},
		}
		zlog.I(`开启HTTPS server：`, m.cfg.TLSAddr)
		go func() { m.srv.ListenAndServeTLS(m.cfg.CertFile, m.cfg.KeyFile) }()

		return

	}

	// 说明都要开启，这里默认将HTTP请求都跳转到HTTPS
	zlog.I(`开启 HTTPS: `, m.cfg.TLSAddr, ` 和 HTTP Server：`, m.cfg.APIAddr)
	m.srv = &http.Server{
		Addr:    m.cfg.TLSAddr,
		Handler: m.router,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			},
		}}

	go func() { m.srv.ListenAndServeTLS(m.cfg.CertFile, m.cfg.KeyFile) }()

	m.sSrv = &http.Server{Addr: m.cfg.APIAddr, Handler: getRedirectFn(m.cfg.TLSAddr)}

	go func() { m.sSrv.ListenAndServe() }()

}

func getRedirectFn(tlsAddr string) http.HandlerFunc {

	_, tlsPort, err := net.SplitHostPort(tlsAddr)

	if err != nil {
		panic(err.Error())
	}
	// tlsPort := tlsAddr

	return func(rw http.ResponseWriter, r *http.Request) {

		host, _, _ := net.SplitHostPort(r.Host)

		target := "https://" + host + ":" + tlsPort + r.URL.Path
		if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		}
		zlog.D("redirect to,", target)
		http.Redirect(rw, r, target,
			// see comments below and consider the codes 308, 302, or 301
			http.StatusTemporaryRedirect)
	}

}
