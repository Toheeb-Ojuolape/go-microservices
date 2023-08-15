package main

import (
	"log"

	"github.com/Toheeb-Ojuolape/go-microservices/internal/database"
	"github.com/Toheeb-Ojuolape/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize database client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
