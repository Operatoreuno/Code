package auth

import (
	"oniplu/db"
	m "oniplu/middleware"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes registra tutte le route di autenticazione su un gruppo/router Gin passato
func AuthRoutes(router gin.IRouter, queries *db.Queries) {
	controller := CreateAuthController(queries)

	router.POST("/signup", m.Decode(controller.SignupHandler))
}
