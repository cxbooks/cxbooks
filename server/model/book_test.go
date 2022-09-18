package model

import (
	"encoding/json"
	"testing"

	"github.com/cxbooks/cxbooks/server/zlog"
	"go.uber.org/zap/zapcore"
)

func initStore() *Store {

	opt := &Opt{
		Driver:   DRSqlite,
		LogLevel: zapcore.DebugLevel,
		Host:     "/data/db/cxbooks.db",
	}
	zlog.I(`打开数据库连接...`)
	store, _ := OpenDB(opt)
	zlog.I(`打开数据库连接`)
	return store
}

func TestRandom(t *testing.T) {
	zlog.Init(`stdout`, zapcore.DebugLevel)
	store := initStore()

	got, err := RandomBooks(store, 10)

	if err != nil {
		t.Fatalf(err.Error())

	}

	for _, book := range got.Books {
		jdata, _ := json.Marshal(&book)

		println(string(jdata))
	}

}

func TestSearchBooks(t *testing.T) {
	zlog.Init(`stdout`, zapcore.DebugLevel)
	req := &Query{
		Search:   "程序",
		PageSize: 10,
		PageNum:  2,
	}

	resp, err := SearchBooks(initStore(), req)

	if err != nil {
		t.Fatalf(err.Error())

	}

	println(`count: `, resp.Total)

	for _, book := range resp.Books {
		jdata, _ := json.Marshal(&book)

		println(string(jdata))
	}

}
