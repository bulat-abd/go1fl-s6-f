package main

import (
	"log"
	"os"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server: ",  log.LstdFlags|log.Lshortfile)
	s := server.NewServer(":8080", logger)
	if err := s.Run(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}