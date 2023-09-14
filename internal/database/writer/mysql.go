package writer

import (
	"database/sql"
	"fmt"
	"log"
	"simple-invitation/configuration"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlWriter struct {
	db *sql.DB
}

type IMysqlWriter interface {
	Close() error
}

func NewMysqlWriter(configDB configuration.IDatabaseWriter) *mysqlWriter {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		strings.TrimSpace(configDB.GetUser()),
		strings.TrimSpace(configDB.GetPassword()),
		strings.TrimSpace(configDB.GetProtocol()),
		strings.TrimSpace(configDB.GetHost()),
		strings.TrimSpace(configDB.GetPort()),
		strings.TrimSpace(configDB.GetName()),
	)
	log.Println("dsnWriter")
	log.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err) // TODO: handle this properly!
		return nil
	}
	return &mysqlWriter{db: db}
}

func (mw *mysqlWriter) GetDB() *sql.DB {
	return mw.db
}

func (mw *mysqlWriter) Close() error {
	return mw.db.Close()
}
