package server

import (
	"log"
	"net/http"
	"time"

	"sprint6_API/internal/handlers"
)

type Server struct {
	Logger   *log.Logger
	MyServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	ServerInstance := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:   logger,
		MyServer: ServerInstance,
	}
}
