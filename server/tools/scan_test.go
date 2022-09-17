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

	store, _ := model.OpenDB(opt)

	return store
}

func TestScanBooks(t *testing.T) {

	zlog.Init(`stdout`, zapcore.DebugLevel)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	manager, _ := NewScannerManager(context.TODO(), `/data/db`, initStore())

	manager.Start(`/data/ebooks`)
	ticker := time.NewTicker(200 * time.Millisecond)

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

	//TODO 清理资源
	zlog.I("收到 ctrl+c 命令....")
	// for d := range data {

	// 	println(d.Path[0])

	// }

}
