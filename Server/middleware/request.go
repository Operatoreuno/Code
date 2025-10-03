package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Aggiunge un Request ID unico ad ogni richiesta.
// Se la richiesta contiene già un header X-Request-Id, lo utilizza,
// altrimenti ne genera uno nuovo.
//
// Il Request ID viene:
//   - Salvato nel contesto Gin con la chiave "request_id"
//   - Aggiunto all'header della risposta X-Request-Id
func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Prova a leggere il Request ID dall'header
		requestID := ctx.GetHeader("X-Request-Id")

		// Se non presente, generane uno nuovo
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Salva nel contesto per uso futuro
		ctx.Set("request_id", requestID)

		// Aggiungi all'header della risposta
		ctx.Header("X-Request-Id", requestID)

		ctx.Next()
	}
}
