package main

import (
	"log"

	"github.com/Equilibriumz/url-shortener/internal/server"
)

func main() {
	s := server.InitServer()

	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
