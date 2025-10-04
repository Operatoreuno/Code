package userauth

import (
	"oniplu/db"
	m "oniplu/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// SetupAuthRoutes registra tutte le route di autenticazione su un gruppo/router Gin passato
func AuthRoutes(router gin.IRouter, queries *db.Queries, redisClient *redis.Client) {
	controller := CreateAuthController(queries, redisClient)

	router.POST("/signup", m.Decode(controller.SignupController))
	router.POST("/login", m.Decode(controller.CookieLoginController))
}
