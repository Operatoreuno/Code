package middleware

import (
	"log"
	"runtime/debug"

	response "oniplu/api/response"
	apiErrors "oniplu/errors"

	"github.com/gin-gonic/gin"
)

// ErrorMiddleware middleware per gestione centralizzata degli errori e panic recovery.
// Gestisce sia i panic che gli errori aggiunti al contesto Gin tramite ctx.Error().
//
// Funzionalità:
//   - Recovery automatico dai panic con stack trace logging
//   - Conversione automatica degli errori in APIError tramite errors.ToAPIError
//   - Logging degli errori server (status 5xx)
//   - Risposta JSON standardizzata tramite response.Error
//
// Note:
//   - Deve essere posizionato dopo RequestIDMiddleware per mantenere il Request ID
//   - Gli errori client (4xx) non vengono loggati, solo gli errori server (5xx)
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v\nStack: %s", err, debug.Stack())

				apiErr := apiErrors.InternalServerError("Si è verificato un errore critico", apiErrors.INTERNAL_SERVER_ERR)
				response.Error(ctx, apiErr)
				ctx.Abort()
			}
		}()

		ctx.Next()

		// Se ci sono errori nel contesto Gin, gestiscili
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			apiErr := apiErrors.ToAPIError(err)

			// Log degli errori server per debug
			if apiErr.ServerError() {
				log.Printf("Server Error: %v", err)
			}

			response.Error(ctx, apiErr)
		}
	}
}
