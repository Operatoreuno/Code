package auth

import (
	"net/http"

	res "oniplu/api/response"
	"oniplu/db"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *AuthService
}

// Crea una nuova istanza di AuthController
func CreateAuthController(queries *db.Queries) *AuthController {
	return &AuthController{
		authService: &AuthService{queries: queries},
	}
}

// SignupHandler gestisce la registrazione di un nuovo utente
func (c *AuthController) SignupHandler(ctx *gin.Context, req *SignupRequest) error {
	// Validation
	if err := ValidateSignupRequest(req); err != nil {
		return err
	}

	// Chiama il service
	user, err := c.authService.Signup(ctx.Request.Context(), req)
	if err != nil {
		return err
	}

	res.Success(ctx, http.StatusOK, "Utente registrato con successo", user)
	return nil
}
