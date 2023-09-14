package configuration

import "os"

type databaseReader struct {
	user     string
	password string
	protocol string
	host     string
	port     string
	name     string
}

type IDatabaseReader interface {
	GetUser() string
	GetPassword() string
	GetProtocol() string
	GetHost() string
	GetPort() string
	GetName() string
}

func NewDatabaseReader() *databaseReader {
	return &databaseReader{
		user:     os.Getenv("DB_USER_READER"),
		password: os.Getenv("DB_PASSWORD_READER"),
		protocol: os.Getenv("DB_PROTOCOL_READER"),
		host:     os.Getenv("DB_HOST_READER"),
		port:     os.Getenv("DB_PORT_READER"),
		name:     os.Getenv("DB_NAME_READER"),
	}
}

func (dw *databaseReader) GetUser() string {
	return dw.user
}

func (dw *databaseReader) GetPassword() string {
	return dw.password
}

func (dw *databaseReader) GetProtocol() string {
	return dw.protocol
}

func (dw *databaseReader) GetHost() string {
	return dw.host
}

func (dw *databaseReader) GetPort() string {
	return dw.port
}

func (dw *databaseReader) GetName() string {
	return dw.name
}
