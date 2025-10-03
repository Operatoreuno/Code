package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	client *redis.Client
}

// NewSessionManager crea un nuovo SessionManager
func SessionManager(client *redis.Client) *Session {
	return &Session{
		client: client,
	}
}

// CreateSession crea una nuova sessione di refresh token in Redis
func (s *Session) CreateSession(ctx context.Context, entityType, entityID, jti string, duration time.Duration) error {
	key := fmt.Sprintf("refresh:%s:%s:%s", entityType, entityID, jti)

	err := s.client.Set(ctx, key, "1", duration).Err()
	if err != nil {
		return fmt.Errorf("errore nella creazione della sessione: %w", err)
	}

	return nil
}

// FindSession cerca una sessione specifica tramite jti in Redis
func (s *Session) FindSession(ctx context.Context, jti string) (bool, error) {
	pattern := fmt.Sprintf("refresh:*:*:%s", jti)

	var keys []string
	iter := s.client.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return false, fmt.Errorf("errore nella ricerca della sessione: %w", err)
	}

	return len(keys) > 0, nil
}

// FindAllSessionByID trova tutte le sessioni per un determinato entityId
func (s *Session) FindAllSessionByID(ctx context.Context, entityType, entityID string) ([]string, error) {
	pattern := fmt.Sprintf("refresh:%s:%s:*", entityType, entityID)

	var sessions []string
	iter := s.client.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		sessions = append(sessions, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("errore nella ricerca delle sessioni: %w", err)
	}

	return sessions, nil
}

// RevokeSession revoca una sessione specifica tramite jti eliminandola da Redis
func (s *Session) RevokeSession(ctx context.Context, jti string) error {
	pattern := fmt.Sprintf("refresh:*:*:%s", jti)

	var keys []string
	iter := s.client.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return fmt.Errorf("errore nella ricerca della sessione: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	err := s.client.Del(ctx, keys...).Err()
	if err != nil {
		return fmt.Errorf("errore nella revoca della sessione: %w", err)
	}

	return nil
}

// RevokeAllSession revoca tutte le sessioni per un determinato entityId
func (s *Session) RevokeAllSession(ctx context.Context, entityType, entityID string) error {
	sessions, err := s.FindAllSessionByID(ctx, entityType, entityID)
	if err != nil {
		return err
	}

	if len(sessions) == 0 {
		return nil
	}

	err = s.client.Del(ctx, sessions...).Err()
	if err != nil {
		return fmt.Errorf("errore nella revoca di tutte le sessioni: %w", err)
	}

	return nil
}

// BlacklistToken aggiunge un access token alla blacklist
func (s *Session) BlacklistToken(ctx context.Context, jti string, duration time.Duration) error {
	key := fmt.Sprintf("blacklist:%s", jti)

	err := s.client.Set(ctx, key, "1", duration).Err()
	if err != nil {
		return fmt.Errorf("errore nell'aggiunta del token alla blacklist: %w", err)
	}

	return nil
}

// FindTokenBlacklisted verifica se un access token è nella blacklist
func (s *Session) FindTokenBlacklisted(ctx context.Context, jti string) (bool, error) {
	key := fmt.Sprintf("blacklist:%s", jti)

	result, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("errore nella verifica della blacklist: %w", err)
	}

	return result > 0, nil
}
