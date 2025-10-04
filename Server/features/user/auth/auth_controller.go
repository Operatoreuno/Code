package userauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"oniplu/auth"
	res "oniplu/api/response"
	"oniplu/db"
)

type AuthController struct {
	authService *AuthService
}

// Crea una nuova istanza di AuthController
func CreateAuthController(queries *db.Queries, redisClient *redis.Client) *AuthController {
	if redisClient == nil {
		panic("Redis client is nil!")
	}
	sessionManager := auth.SessionManager(redisClient)
	if sessionManager == nil {
		panic("Session manager is nil!")
	}
	return &AuthController{
		authService: &AuthService{
			queries: queries,
			session: sessionManager,
		},
	}
}

func (c *AuthController) CheckUserController(ctx *gin.Context, req *CheckUserRequest) error {
	return nil
}

func (c *AuthController) SignupController(ctx *gin.Context, req *SignupRequest) error {
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

func (c *AuthController) UserMeController(ctx *gin.Context) error {
	return nil
}

//Web

func (c *AuthController) CookieLoginController(ctx *gin.Context, req *LoginRequest) error {
	// Validation
	if err := ValidateLoginRequest(req); err != nil {
		return err
	}

	login, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		return err
	}

	// Imposta il refresh token come cookie httpOnly
	config := auth.AuthUserConfig
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie(
		"refreshToken",         // name
		login.RefreshToken,     // value
		config.RefreshDuration, // maxAge in secondi
		"/",                    // path
		"",                     // domain (vuoto = dominio corrente)
		false,                  // secure (false per sviluppo, true in produzione per HTTPS)
		true,                   // httpOnly
	)

	// Restituisci solo l'access token nel body
	res.Success(ctx, http.StatusOK, "Login effettuato con successo", gin.H{
		"accessToken": login.AccessToken,
		"user":        login.User,
	})
	return nil
}

func (c *AuthController) CookieLogoutController(ctx *gin.Context) error {
	return nil
}

func (c *AuthController) CookieRefreshController(ctx *gin.Context) error {
	return nil
}

//Mobile

func (c *AuthController) LoginController(ctx *gin.Context, req *LoginRequest) error {
	// Validation
	if err := ValidateLoginRequest(req); err != nil {
		return err
	}

	login, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		return err
	}

	// Restituisci access token e refresh token nel body (per mobile)
	res.Success(ctx, http.StatusOK, "Login effettuato con successo", login)
	return nil
}

func (c *AuthController) LogoutController(ctx *gin.Context) error {
	return nil
}

func (c *AuthController) RefreshController(ctx *gin.Context) error {
	return nil
}
