package main

import (
	"log"

	"oniplu/api"
	"oniplu/config"
	"oniplu/db"
	"oniplu/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connessione al database
	database, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Errore nella connessione al database: %v", err)
	}
	defer config.CloseDatabase(database)

	// Crea istanza singleton delle queries
	queries := db.Create(database)

	// Inizializza Gin
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Logger(), gin.Recovery())

	// Middleware (ordine importante)
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.SecurityMiddleware())
	router.Use(middleware.ErrorMiddleware())

	// Prefisso API
	apiGroup := router.Group("/api")
	api.APIRoutes(apiGroup, queries)

	// Avvia il server con graceful shutdown
	config.StartServer(router)
}
