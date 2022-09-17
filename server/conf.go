package server

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/jinzhu/configor"
	"go.uber.org/zap/zapcore"
)

// 变量
var (
	VERSION    = "UNKNOWN"
	buildstamp = "UNKNOWN"
	githash    = "UNKNOWN"

	ConfFile    string
	showVersion bool
)

// Config 定义 配置结构图
type Config struct {
	APIAddr  string `yaml:"api_addr"`
	TLSAddr  string `yaml:"tls_addr"` //是否开启https
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
	DataPath string `yaml:"data_path"`

	DBOpt *model.Opt `yaml:"db_opt" json:"db_opt"`

	LogLevel zapcore.Level `yaml:"verbose"`
	LogDir   string        `yaml:"log_dir"`

	InitFlag bool `yaml:"init_flag"`

	PPROF string `yaml:"pprof"` //模块启动模式，主模式加载所有路由，从模式仅仅加载自身路由
}

// InitConfig 初始化
func InitConfig() *Config {

	config := &Config{}
	flag.BoolVar(&showVersion, "version", false, "show build version.")

	flag.StringVar(&ConfFile, "conf", "conf/conf.yml", "The configure file")
	flag.StringVar(&config.PPROF, "pprof", "", "[localhost:6060]start debug page.")
	flag.StringVar(&config.APIAddr, "addr", "", "The api listen addr[192.168.1.201:8000] if not set this config read from conf.yml .")
	flag.StringVar(&config.LogDir, "log", "stdout", "The log file")
	flag.Var(&config.LogLevel, "verbose", "The log level [debug,info,error]")
	flag.BoolVar(&config.InitFlag, "init", false, "init db.")
	flag.Parse()
	changeWorkspace()

	if showVersion {
		println(`core version: `, VERSION)
		println(`Git Commit Hash: `, githash)
		println(`UTC Build Time: `, buildstamp)
		os.Exit(0)
	}

	logDir := config.LogDir
	if logDir != `stdout` {
		logDir, _ = filepath.Abs(config.LogDir)
	}

	zlog.Init(logDir, config.LogLevel)

	defer zlog.Flush()

	zlog.I("当前版本: ", VERSION)
	zlog.I(`Git Commit Hash: `, githash)
	zlog.I(`UTC Build Time: `, buildstamp)
	zlog.I(`当前日志等级为: `, config.LogLevel.CapitalString())
	ConfFile, _ = filepath.Abs(ConfFile)

	if _, err := os.Stat(ConfFile); os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.E(`配置文件路径不存在: `, ConfFile)
		os.Exit(1)
	}

	if err := configor.Load(config, ConfFile); err != nil {
		zlog.E(`加载配置文件失败: `, err.Error())
		os.Exit(1)
	}

	if config.APIAddr == "" && config.TLSAddr == "" { //监听地址必须只是配置一个
		zlog.E(`API监听地址格式异常`)
		os.Exit(2)
	}

	if config.TLSAddr != "" { //设置了TLS 则校验 证书与秘钥路径是否存在

		if _, err := os.Stat(config.CertFile); os.IsNotExist(err) {
			// path/to/whatever does not exist
			zlog.E(`证书路径不存在: `, config.CertFile)
			os.Exit(1)
		}
		if _, err := os.Stat(config.KeyFile); os.IsNotExist(err) {
			// path/to/whatever does not exist
			zlog.E(`秘钥路径不存在: `, config.KeyFile)
			os.Exit(1)
		}
	}

	if config.DBOpt != nil {
		config.DBOpt.LogLevel = config.LogLevel
	}

	if config.DataPath == "" { //监听地址必须只是配置一个
		config.DataPath = `/data/cache/`
	}

	return config
}

// changeWorkspace 修改当前程序workspace
func changeWorkspace() {
	//设置当前工作目录
	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := filepath.Base(binDir)
	workDir := strings.TrimRight(strings.TrimRight(binDir, file), "/bin")
	os.Chdir(workDir)
	// log.D(`修改当前工作路径：`, workDir)
}
