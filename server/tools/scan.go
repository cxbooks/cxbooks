// scan 目录扫描搜刮工具
package tools

import (
	"context"
	"time"
)

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
