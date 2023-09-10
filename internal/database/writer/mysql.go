package writer

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlWriter struct {
	db *sql.DB
}

type IMysqlWriter interface {
	Close() error
}

func NewMysqlWriter() *mysqlWriter {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		strings.TrimSpace(os.Getenv("DB_USER_WRITER")),
		strings.TrimSpace(os.Getenv("DB_PASSWORD_WRITER")),
		strings.TrimSpace(os.Getenv("DB_PROTOCOL_WRITER")),
		strings.TrimSpace(os.Getenv("DB_HOST_WRITER")),
		strings.TrimSpace(os.Getenv("DB_PORT_WRITER")),
		strings.TrimSpace(os.Getenv("DB_NAME_WRITER")),
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
