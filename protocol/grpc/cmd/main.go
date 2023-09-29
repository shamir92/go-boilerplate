package main

import (
	"fmt"
	"log"
	"net"
	configuration "simple-invitation/configuration"
	reader "simple-invitation/internal/database/reader"
	writer "simple-invitation/internal/database/writer"
	"simple-invitation/protocol/grpc/controllers"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("TCP Not Up")
	}

	server := grpc.NewServer()

	// For configuration
	var dbReaderConfiguration configuration.IDatabaseReader = configuration.NewDatabaseReader()
	var dbWriterConfiguration configuration.IDatabaseWriter = configuration.NewDatabaseWriter()
	// For External Interfaces
	reader := reader.NewMysqlReader(dbReaderConfiguration)
	writer := writer.NewMysqlWriter(dbWriterConfiguration)

	var memberController controllers.IMemberController = controllers.NewMemberController(reader.GetDB(), writer.GetDB(), server)
	log.Println(memberController)
	err = server.Serve(list)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}

	// routes.PublicRoute(r, reader.GetDB(), writer.GetDB())

}
