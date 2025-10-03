package res

import (
	apiErrors "oniplu/errors"

	"github.com/gin-gonic/gin"
)

// Struttura standard per le risposte API
type APIResponse struct {
	Success   bool                 `json:"success"`
	Message   string               `json:"message"`
	ErrorCode *apiErrors.ErrorCode `json:"errorCode,omitempty"`
	Data      any                  `json:"data"`
}

// Risposta di successo
func Success(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Risposta di errore (accetta APIError strutturato)
func Error(ctx *gin.Context, err *apiErrors.APIError) {
	ctx.JSON(err.StatusCode, APIResponse{
		Success:   false,
		Message:   err.Message,
		ErrorCode: &err.ErrorCode,
		Data:      err.Data,
	})
}
