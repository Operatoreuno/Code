package auth

// TokenPayload rappresenta il payload standardizzato per i token JWT
type TokenPayload struct {
	EntityID string `json:"entity_id"` // entityId (userId o adminId)
	Jti      string `json:"jti"`       // unique id per token (blacklist/refresh store)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
