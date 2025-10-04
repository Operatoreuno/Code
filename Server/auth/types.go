package auth

// TokenPayload rappresenta il payload standardizzato per i token JWT
type TokenPayload struct {
	EntityID string // entityId (userId o adminId)
	Jti      string // unique id per token (blacklist/refresh store)
}

type AuthConfig struct {
	AccessSecret    string
	RefreshSecret   string
	AccessDuration  int
	RefreshDuration int
}

func UserAuthConfig() AuthConfig {
	return AuthConfig{
		AccessSecret:    "Ciao",
		RefreshSecret:   "Ciao",
		AccessDuration:  7 * 24 * 60 * 60,
		RefreshDuration: 30 * 60,
	}
}

var AuthUserConfig = UserAuthConfig()

func AdminAuthConfig() AuthConfig {
	return AuthConfig{
		AccessSecret:    "Ciao",
		RefreshSecret:   "Ciao",
		AccessDuration:  15 * 60,
		RefreshDuration: 48 * 60 * 60,
	}
}

var AuthAdminConfig = AdminAuthConfig()
