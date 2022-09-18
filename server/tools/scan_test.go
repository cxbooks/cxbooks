package tools

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
	"go.uber.org/zap/zapcore"
)

func initStore() *model.Store {

	opt := &model.Opt{
		Driver:   model.DRSqlite,
		LogLevel: zapcore.DebugLevel,
		Host:     "/data/db/cxbooks.db",
	}
	zlog.I(`打开数据库连接...`)
	store, _ := model.OpenDB(opt)
	zlog.I(`打开数据库连接`)
	return store
}

func TestScanBooks(t *testing.T) {

	zlog.Init(`stdout`, zapcore.DebugLevel)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	manager, _ := NewScannerManager(context.TODO(), `/data/db`, initStore())

	manager.Start(`/data/ebooks`, 1)
	ticker := time.NewTicker(3 * time.Second)

	for {

		select {
		case <-ch:
			return
		case <-ticker.C:
			states := manager.Status()

			str, _ := json.Marshal(states)

			println(string(str))

		}

	}

}
