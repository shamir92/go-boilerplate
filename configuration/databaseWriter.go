package configuration

import "os"

type databaseWriter struct {
	user     string
	password string
	protocol string
	host     string
	port     string
	name     string
}

type IDatabaseWriter interface {
	GetUser() string
	GetPassword() string
	GetProtocol() string
	GetHost() string
	GetPort() string
	GetName() string
}

func NewDatabaseWriter() *databaseWriter {
	return &databaseWriter{
		user:     os.Getenv("DB_USER_WRITER"),
		password: os.Getenv("DB_PASSWORD_WRITER"),
		protocol: os.Getenv("DB_PROTOCOL_WRITER"),
		host:     os.Getenv("DB_HOST_WRITER"),
		port:     os.Getenv("DB_PORT_WRITER"),
		name:     os.Getenv("DB_NAME_WRITER"),
	}
}

func (dw *databaseWriter) GetUser() string {
	return dw.user
}

func (dw *databaseWriter) GetPassword() string {
	return dw.password
}

func (dw *databaseWriter) GetProtocol() string {
	return dw.protocol
}

func (dw *databaseWriter) GetHost() string {
	return dw.host
}

func (dw *databaseWriter) GetPort() string {
	return dw.port
}

func (dw *databaseWriter) GetName() string {
	return dw.name
}
