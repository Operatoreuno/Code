package api

import (
	"oniplu/db"
	"oniplu/features/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// SetupAPIRoutes registra tutte le route API su un gruppo/router Gin passato
func APIRoutes(router gin.IRouter, queries *db.Queries, redisClient *redis.Client) {
	// Gruppo principale: /user
	userGroup := router.Group("/user")
	user.UserRoutes(userGroup, queries, redisClient)

	// adminGroup := router.Group("/admin")
	// admin.SetupAdminRoutes(adminGroup, db)

	// sseGroup := router.Group("/sse")
	// sse.SetupSSERoutes(sseGroup, db)

	// stripeGroup := router.Group("/stripe")
	// stripe.SetupStripeRoutes(stripeGroup, db)
}
