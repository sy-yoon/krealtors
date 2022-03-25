package gms

type Configuration interface {
	GetBindPort() int
	GetDBType() string
	GetDBUser() string
	GetDBPassword() string
	GetDBHost() string
	GetDBPort() int
	GetDatabase() string
	GetDBMaxIdelConns() int
	GetDBMaxOpenConns() int
	GetLogLevel() string
}
