package main

import (
	"log"
	"os"

	"go-pipeline/initializers"
)

func main() {
	db, err := initializers.ConnectDatabase()
	if err != nil {
		log.Fatalf("initialize database: %v", err)
	}
	app := initializers.NewService(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("start HTTP server: %v", err)
	}
}
