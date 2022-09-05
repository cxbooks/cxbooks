package model

import (
	"strconv"

	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Driver 定义数据库类型
type Driver string

const (
	// DRPostgres driver name for pg
	DRPostgres Driver = "postgres"
	// DRMySQL driver name for mysql
	DRMySQL Driver = "mysql"
	// DRSqlite driver name for mysql
	DRSqlite Driver = "sqlite"
)

// Opt PG 数据库配置
type Opt struct {
	Driver   Driver        `yaml:"driver" json:"driver"`       //数据库类型
	LogLevel zapcore.Level `yaml:"log_level" json:"log_level"` //日志等级
	Host     string        `yaml:"host" json:"host"`           //数据地址
	Port     int           `yaml:"port" json:"port"`           //端口
	User     string        `yaml:"user" json:"user"`           //用户名
	DBName   string        `yaml:"dbname" json:"dbname"`       //数据库名
	Passwd   string        `yaml:"passwd" json:"passwd"`       //密码
	SSLMode  string        `yaml:"sslmode" json:"sslmode"`     // disable, varify-full //SSL选项
	Args     string        `yaml:"args" json:"args"`           // charset=utf8 //额外选项
}

// DSN return gorm v2 Dialector
func (opt *Opt) DSN() gorm.Dialector {
	switch {
	case opt.Driver == `postgres`:
		return opt.PGDSN()
	case opt.Driver == `mysql`:
		return opt.MySQLDSN()
	case opt.Driver == `sqlite`:
		return opt.SQLiteDSN()
	}

	return opt.PGDSN()
}

// PGDSN  转换为 PG 连接字符串
//
//	postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca
//
// postgres://jack:secret@foo.example.com:5432,bar.example.com:5432/mydb
func (opt *Opt) PGDSN() gorm.Dialector {

	// dsn := "host=" + opt.Host + " port=" + strconv.Itoa(opt.Port) + " user=" + opt.User + " dbname=" + opt.DBName + " password=" + opt.Passwd + " sslmode=" + opt.SSLMode
	dsn := `postgres://` + opt.User + `:` + opt.Passwd + `@` + opt.Host

	if opt.Port != 0 { //这里默认如果端口未么有添加写认为端口地址存储在HOST目录
		dsn += `:` + strconv.Itoa(opt.Port)
	}

	if opt.SSLMode == `` {
		opt.SSLMode = `disable`
	}

	dsn += `/` + opt.DBName + `?sslmode=` + opt.SSLMode + `&` + opt.Args

	return postgres.Open(dsn)
}

// MySQLDSN  转换为 mysql 连接字符串
func (opt *Opt) MySQLDSN() gorm.Dialector {

	if opt.Args == "" {
		opt.Args = `charset=utf8&parseTime=True&loc=Local`
	}

	dsn := opt.User + `:` + opt.Passwd + `@tcp(` + opt.Host + ":" + strconv.Itoa(opt.Port) + ")/" + opt.DBName + "?" + opt.Args
	return mysql.Open(dsn)
}

// SQLiteDSN  返回 sqlite 数据库连接字符
func (opt *Opt) SQLiteDSN() gorm.Dialector {

	return sqlite.Open(opt.Host)
}

func (opt *Opt) String() string {
	return "driver=" + string(opt.Driver) + " host=" + opt.Host + " port=" + strconv.Itoa(opt.Port) + " user=" + opt.User + " dbname=" + opt.DBName + " password=xxxxxx sslmode=" + opt.SSLMode
}