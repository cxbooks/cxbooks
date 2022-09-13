package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/cxbooks/cxbooks/server"

	"github.com/cxbooks/cxbooks/server/zlog"
)

func main() {

	api := server.NewService()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	api.StartContext(ctx)

	<-ch
	//TODO 清理资源
	zlog.I("收到 ctrl+c 命令....")
	defer cancel()
	api.GracefulStop()

}
