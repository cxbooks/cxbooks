package model

import (
	"context"
	"time"

	"github.com/cxbooks/cxbooks/zlog"
	"gorm.io/gorm"
)

// OpenDB 直接打开数据库连接 如果失败立即返回错误
func OpenDB(opt *Opt) (*gorm.DB, error) {

	db, err := gorm.Open(opt.DSN(), &gorm.Config{})
	if err != nil {
		// log.E(`链接数据库失败：`,, err.Error())
		zlog.I(`链接数据库异常:`, opt.String())
	}

	return db, err

}

// WaitDB 打开数据库连接，如果失败则一直尝试重连直到成功为止
func WaitDB(ctx context.Context, opt *Opt) (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	f := func() error {

		db, err = gorm.Open(opt.DSN(), &gorm.Config{})
		if err != nil {
			zlog.I(`链接数据库异常:`, opt.String())
		}
		return err

	}

	err = Wait(ctx, 30, f)

	return db, err
}

// Wait 等待某个函数执行成功
func Wait(ctx context.Context, sleep time.Duration, f func() error) error {

	if err := f(); err == nil {
		return nil
	}
	for {

		ticker := time.NewTicker(time.Second * sleep)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := f(); err == nil {
				return nil
			}

		}
	}

}
