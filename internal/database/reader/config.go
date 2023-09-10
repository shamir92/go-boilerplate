package reader

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// belum tahu mau gunakan gmana
type configurationReader struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func initDatabaseReader() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST_READER"),
		os.Getenv("DB_USER_READERR"),
		os.Getenv("DB_PASSWORD_READER"),
		os.Getenv("DB_NAME_READER"),
		os.Getenv("DB_PORT_READER"),
		os.Getenv("DB_SSL_MODE_READER"),
		os.Getenv("DB_TIMEZONE_READER"),
	)
	dbReader, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	dbReader.SetConnMaxLifetime(time.Minute * 3)
	dbReader.SetMaxOpenConns(10)
	dbReader.SetMaxIdleConns(10)
}
