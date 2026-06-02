package main

import (
	"log"
	"os"
	"sprint6_API/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Starting server at port 8080")

	err := srv.MyServer.ListenAndServe()
	if err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}
