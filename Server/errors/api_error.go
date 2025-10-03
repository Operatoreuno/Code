package apiErrors

import (
	"fmt"
)

// APIError rappresenta un errore API strutturato con codice, messaggio e dati opzionali
type APIError struct {
	StatusCode int       `json:"-"`
	ErrorCode  ErrorCode `json:"errorCode"`
	Message    string    `json:"message"`
	Data       any       `json:"data,omitempty"`
}

// Error implementa l'interfaccia error standard di Go
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s (Code: %d)", e.StatusCode, e.Message, e.ErrorCode)
}

// ClientError verifica se l'errore è un errore client (4xx)
func (e *APIError) ClientError() bool {
	return e.StatusCode >= 400 && e.StatusCode < 500
}

// ServerError verifica se l'errore è un errore server (5xx)
func (e *APIError) ServerError() bool {
	return e.StatusCode >= 500
}

// Converte errori generici in APIError per fornire una risposta API coerente
func ToAPIError(err error) *APIError {

	if apiErr, ok := err.(*APIError); ok {
		return apiErr
	}
	return InternalServerError("Si è verificato un errore imprevisto", INTERNAL_SERVER_ERR)
}
