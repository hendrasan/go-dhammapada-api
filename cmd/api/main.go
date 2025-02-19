package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hendrasan/go-dhammapada-api/config"
	"github.com/hendrasan/go-dhammapada-api/internal/handlers"
	"github.com/hendrasan/go-dhammapada-api/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	r := gin.Default()

	handlers.RegisterRoutes(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
