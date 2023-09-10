package reader

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlReader struct {
	db *sql.DB
}

type IMysqlReader interface {
	Close() error
}

func NewMysqlReader() *mysqlReader {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		strings.TrimSpace(os.Getenv("DB_USER_READER")),
		strings.TrimSpace(os.Getenv("DB_PASSWORD_READER")),
		strings.TrimSpace(os.Getenv("DB_PROTOCOL_READER")),
		strings.TrimSpace(os.Getenv("DB_HOST_READER")),
		strings.TrimSpace(os.Getenv("DB_PORT_READER")),
		strings.TrimSpace(os.Getenv("DB_NAME_READER")),
	)
	log.Println("dsnReader")
	log.Println(dsn)

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
