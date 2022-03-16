package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	//Mysql MySql
	Sqlite Sqlite
}
type Sqlite struct {
	Dsn string
}

type MySql struct {
	Username     string
	Password     string
	Hostname     string
	Hostport     string
	Database     string
	MaxOpenConns int
	MaxIdleConns int
}
