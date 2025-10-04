package user

import (
	"oniplu/db"
	userAuth "oniplu/features/user/auth"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// SetupUserRoutes registra tutte le route della feature user su un gruppo/router Gin passato
func UserRoutes(router gin.IRouter, queries *db.Queries, redisClient *redis.Client) {
	// Sottogruppo: /auth
	authGroup := router.Group("/auth")
	userAuth.AuthRoutes(authGroup, queries, redisClient)

	// In futuro potrai aggiungere altre sub-feature:
	// profileGroup := router.Group("/profile")
	// profile.ProfileRoutes(profileGroup, db)
	// settingsGroup := router.Group("/settings")
	// settings.SetupSettingsRoutes(settingsGroup, db)
}
