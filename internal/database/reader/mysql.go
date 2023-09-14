package reader

import (
	"database/sql"
	"fmt"
	"log"
	"simple-invitation/configuration"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlReader struct {
	db *sql.DB
}

type IMysqlReader interface {
	Close() error
}

func NewMysqlReader(configDB configuration.IDatabaseReader) *mysqlReader {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		strings.TrimSpace(configDB.GetUser()),
		strings.TrimSpace(configDB.GetPassword()),
		strings.TrimSpace(configDB.GetProtocol()),
		strings.TrimSpace(configDB.GetHost()),
		strings.TrimSpace(configDB.GetPort()),
		strings.TrimSpace(configDB.GetName()),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err) // TODO: handle this properly!
		return nil
	}
	return &mysqlReader{db: db}
}

func (mr *mysqlReader) GetDB() *sql.DB {
	return mr.db
}

func (mr *mysqlReader) Close() error {
	return mr.db.Close()
}
