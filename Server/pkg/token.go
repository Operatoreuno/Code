package pkg

import (
	"crypto/sha256"
	"fmt"
	"oniplu/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken genera un token di accesso JWT con durata limitata
func GenerateToken(payload auth.TokenPayload, secret string, duration time.Duration) (string, error) {

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": payload.EntityID,
		"jti": payload.Jti,
		"iat": now.Unix(),
		"exp": now.Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// VerifyToken verifica e decodifica un token JWT utilizzando il segreto fornito
func VerifyToken(tokenString, secret string) (*auth.TokenPayload, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodo di firma non atteso: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token non valido: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token non valido")
	}

	return &auth.TokenPayload{
		EntityID: claims.Subject,
		Jti:      claims.ID,
	}, nil
}

// HashToken genera un hash SHA-256 di un token per il salvataggio sicuro nel database
func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", hash)
}

// ExtractTokenFromHeader estrae il token di accesso dall'header Authorization
func ExtractAccessToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header mancante")
	}

	const prefix = "Bearer "

	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		return "", fmt.Errorf("formato authorization header non valido")
	}

	return authHeader[len(prefix):], nil
}

// ExtractRefreshToken estrae il refresh token provando prima dai cookie, poi dall'header
func ExtractRefreshToken(cookies map[string]string, headerValue string) (string, error) {
	// Prova prima dai cookie
	if token, exists := cookies["refreshToken"]; exists && token != "" {
		return token, nil
	}

	// Se non trovato nei cookie, prova dall'header
	if headerValue != "" {
		return headerValue, nil
	}

	// Se non trovato in nessuno dei due
	return "", fmt.Errorf("refresh token mancante")
}
