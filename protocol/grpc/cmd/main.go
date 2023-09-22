package main

import (
	"log"
	configuration "simple-invitation/configuration"
	reader "simple-invitation/internal/database/reader"
	writer "simple-invitation/internal/database/writer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// For configuration
	var dbReaderConfiguration configuration.IDatabaseReader = configuration.NewDatabaseReader()
	var dbWriterConfiguration configuration.IDatabaseWriter = configuration.NewDatabaseWriter()

	// For External Interfaces
	reader := reader.NewMysqlReader(dbReaderConfiguration)
	writer := writer.NewMysqlWriter(dbWriterConfiguration)
	log.Println(reader, writer)

	// routes.PublicRoute(r, reader.GetDB(), writer.GetDB())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}