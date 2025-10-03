package auth

import (
	"context"
	"database/sql"
	"oniplu/db"
	apiErrors "oniplu/errors"
	"oniplu/pkg"
)

type AuthService struct {
	queries *db.Queries
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
	hashedPassword, err := pkg.HashPasswordWithArgon2(req.Password)

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
