package auth

import (
	apiError "oniplu/errors"
	"oniplu/pkg"
)

// ValidationDetail rappresenta i dettagli di un errore di validazione
type ValidationDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Valida i dati di registrazione
func ValidateSignupRequest(req *SignupRequest) error {
	errorDetails := make([]ValidationDetail, 0)

	// Valida email (required)
	if err := pkg.EmailValidator(req.Email, true); err != nil {
		errorDetails = append(errorDetails, ValidationDetail{
			Field:   "email",
			Message: err.Error(),
		})
	}

	// Valida password (required)
	if err := pkg.PasswordValidator(req.Password, true); err != nil {
		errorDetails = append(errorDetails, ValidationDetail{
			Field:   "password",
			Message: err.Error(),
		})
	}

	// Valida name (required, 1-50 caratteri)
	if err := pkg.StringValidator(req.Name, 1, 50, true); err != nil {
		errorDetails = append(errorDetails, ValidationDetail{
			Field:   "name",
			Message: err.Error(),
		})
	}

	// Valida surname (required, 1-50 caratteri)
	if err := pkg.StringValidator(req.Surname, 1, 50, true); err != nil {
		errorDetails = append(errorDetails, ValidationDetail{
			Field:   "surname",
			Message: err.Error(),
		})
	}

	if len(errorDetails) > 0 {
		return apiError.BadRequestError("Dati di input non validi", apiError.INVALID_REQUEST, errorDetails)
	}
	return nil
}
