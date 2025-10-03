package apiErrors

type ErrorCode int

const (
	// Errori generici 1000-1999
	NOT_FOUND            ErrorCode = 1000
	ALREADY_EXISTS       ErrorCode = 1001
	INVALID_REQUEST      ErrorCode = 1002
	INTERNAL_SERVER_ERR  ErrorCode = 1004
	SERVICE_UNAVAILABLE  ErrorCode = 1005
	CONSTRAINT_VIOLATION ErrorCode = 1006
	TOMANY_REQUESTS      ErrorCode = 1007
	CONFLICT             ErrorCode = 1008
	UNPROCESSABLE        ErrorCode = 1009

	// Errori di autenticazione/autorizzazione 4000-4999
	UNAUTHORIZED ErrorCode = 4000
	FORBIDDEN    ErrorCode = 4010

	// Errori token 4001-4009
	TOKEN_MISSING  ErrorCode = 4001
	TOKEN_EXPIRED  ErrorCode = 4002
	TOKEN_REVOKED  ErrorCode = 4003
	TOKEN_INVALID  ErrorCode = 4004
	TOKEN_MISMATCH ErrorCode = 4005
)
