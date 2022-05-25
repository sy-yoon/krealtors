package config

import (
	_ "github.com/go-sql-driver/mysql"
)

var Settings ServerConfig

type ServerConfig struct {
	DatabaseType   string `mapstructure:"DB_TYPE"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	Database       string `mapstructure:"DATABASE"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         int    `mapstructure:"DB_PORT"`
	DBMaxIdelConns int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBMaxOpenConns int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	BindAddress    string `mapstructure:"BIND_ADDRESS"`
	BindPort       int    `mapstructure:"BIND_PORT"`
	StaticContents string `mapstructure:"STATIC_CONTENTS_PATH"`
}

func (me *ServerConfig) GetBindPort() int {
	return me.BindPort
}

func (me *ServerConfig) GetDBType() string {
	return me.DatabaseType
}

func (me *ServerConfig) GetDBHost() string {
	return me.DBHost
}

func (me *ServerConfig) GetDBPort() int {
	return me.DBPort
}

func (me *ServerConfig) GetDatabase() string {
	return me.Database
}

func (me *ServerConfig) GetDBPassword() string {
	return me.DBPassword
}

func (me *ServerConfig) GetDBUser() string {
	return me.DBUser
}

func (me *ServerConfig) GetDBMaxIdelConns() int {
	return me.DBMaxIdelConns
}

func (me *ServerConfig) GetDBMaxOpenConns() int {
	return me.DBMaxOpenConns
}

func (me *ServerConfig) GetLogLevel() string {
	return me.LogLevel
}
