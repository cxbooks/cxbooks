package tools

import (
	"testing"

	"github.com/cxbooks/cxbooks/server/zlog"
	"go.uber.org/zap/zapcore"
)

func TestScanBooks(t *testing.T) {

	zlog.Init(`stdout`, zapcore.DebugLevel)

	path := `/data/ebooks`

	data := Scan(path, 10)

	for d := range data {

		println(d.Path[0])

	}

}
