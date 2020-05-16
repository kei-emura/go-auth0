package main

import (
	"go-auth0/app"
	"go-auth0/server"
	"log"
)

func main() {
	app.Init()
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
