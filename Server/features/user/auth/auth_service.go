package userauth

import (
	"context"
	"database/sql"
	"time"

	"oniplu/auth"
	"oniplu/db"
	apiErrors "oniplu/errors"

	"github.com/google/uuid"
)

type AuthService struct {
	queries *db.Queries
	session *auth.Session
}

func (s *AuthService) Signup(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	// Verifica se l'email è già in uso usando sqlc
	count, err := s.queries.CheckUserEmailExists(ctx, req.Email)

	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nella verifica email", apiErrors.INTERNAL_SERVER_ERR)
	}

	if count > 0 {
		return nil, apiErrors.ConflictError("Email già in uso", apiErrors.ALREADY_EXISTS)
	}

	// Hash della password con Argon2
	hashedPassword, err := auth.HashPassword(req.Password)

	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nell'hashing della password", apiErrors.INTERNAL_SERVER_ERR)
	}

	// TODO: Sostituire con vera chiamata a Stripe API
	stripeCustomerID := sql.NullString{String: "teststripeid", Valid: true}

	// Inserisci l'utente nel database con stripe_customer_id temporaneo
	user, err := s.queries.CreateUser(
		ctx,
		req.Email,
		hashedPassword,
		req.Name,
		req.Surname,
		stripeCustomerID,
	)
	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nella creazione utente", apiErrors.INTERNAL_SERVER_ERR)
	}

	return &SignupResponse{
		ID:               user.ID.String(),
		Email:            user.Email,
		Name:             user.Name,
		Surname:          user.Surname,
		StripeCustomerId: user.StripeCustomerID.String,
	}, nil
}

func (s *AuthService) CheckUser(ctx context.Context, req *SignupRequest) {
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// 1. Trova utente by email
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apiErrors.NotFoundError("Utente non trovato", apiErrors.NOT_FOUND)
		}
		return nil, apiErrors.InternalServerError("Errore nella ricerca utente", apiErrors.INTERNAL_SERVER_ERR)
	}

	// 2. Verifica la password inserita con quella del DB
	isValid, err := auth.VerifyPassword(req.Password, user.Password)
	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nella verifica password", apiErrors.INTERNAL_SERVER_ERR)
	}

	if !isValid {
		return nil, apiErrors.BadRequestError("Credenziali errate", apiErrors.INVALID_REQUEST, nil)
	}

	// 3. Crea il payload
	jti := uuid.New().String()
	payload := auth.TokenPayload{
		EntityID: user.ID.String(),
		Jti:      jti,
	}

	// 4. Genera access token e refresh token seguendo la config utente
	config := auth.AuthUserConfig
	accessToken, err := auth.GenerateToken(
		payload,
		config.AccessSecret,
		time.Duration(config.AccessDuration)*time.Second,
	)
	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nella generazione access token", apiErrors.INTERNAL_SERVER_ERR)
	}

	refreshToken, err := auth.GenerateToken(
		payload,
		config.RefreshSecret,
		time.Duration(config.RefreshDuration)*time.Second,
	)
	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nella generazione refresh token", apiErrors.INTERNAL_SERVER_ERR)
	}

	if s.session == nil {
		return nil, apiErrors.InternalServerError("Session manager non inizializzato", apiErrors.INTERNAL_SERVER_ERR)
	}

	err = s.session.RollSession(
		ctx,
		"user",
		user.ID.String(),
		jti,
		time.Duration(config.RefreshDuration)*time.Second,
		2,
	)
	if err != nil {
		return nil, apiErrors.InternalServerError("Errore nella gestione sessioni", apiErrors.INTERNAL_SERVER_ERR)
	}

	// 6. Restituisci LoginResponse
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: User{
			ID:       user.ID.String(),
			Email:    user.Email,
			Name:     user.Name,
			Surname:  user.Surname,
			Phone:    user.Phone.String,
			IsActive: user.IsActive,
		},
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, req *TokenRequest) {
}

func (s *AuthService) Logout(ctx context.Context, req *TokenRequest) {
}
