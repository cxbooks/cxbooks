package safe

import (
	"runtime/debug"

	"github.com/cxbooks/cxbooks/server/zlog"
)

// Go starts a recoverable goroutine.
func Go(goroutine func()) {
	GoWithRecover(goroutine, defaultRecoverGoroutine)
}

// GoWithRecover starts a recoverable goroutine using given customRecover() function.
func GoWithRecover(goroutine func(), customRecover func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				customRecover(err)
			}
		}()
		goroutine()
	}()
}

func defaultRecoverGoroutine(err interface{}) {

	zlog.E("Error in Go routine: ", err)
	zlog.E("Stack: ", debug.Stack())
}
