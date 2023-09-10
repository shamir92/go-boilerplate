package main

import (
	"simple-invitation/api/routes"
	reader "simple-invitation/internal/database/reader"
	writer "simple-invitation/internal/database/writer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	reader := reader.NewMysqlReader()
	writer := writer.NewMysqlWriter()
	routes.PublicRoute(r, reader.GetDB(), writer.GetDB())

	// register repository
	// memberRepository := repository.NewMemberRepository(reader.GetDB(), writer.GetDB())
	// log.Println(memberRepository)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
