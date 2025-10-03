package user

import (
	"oniplu/db"
	"oniplu/features/user/auth"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes registra tutte le route della feature user su un gruppo/router Gin passato
func UserRoutes(router gin.IRouter, queries *db.Queries) {
	// Sottogruppo: /auth
	authGroup := router.Group("/auth")
	auth.AuthRoutes(authGroup, queries)

	// In futuro potrai aggiungere altre sub-feature:
	// profileGroup := router.Group("/profile")
	// profile.ProfileRoutes(profileGroup, db)
	// settingsGroup := router.Group("/settings")
	// settings.SetupSettingsRoutes(settingsGroup, db)
}
