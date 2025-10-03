package config

import (
	"context"
	"fmt"
	"log"
	"oniplu/pkg"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisURL restituisce l'URL di Redis dalle variabili d'ambiente
func RedisURL() string {
	return pkg.GetEnv("REDIS_URL", "localhost:6379")
}

// RedisPassword restituisce la password di Redis dalle variabili d'ambiente
func RedisPassword() string {
	return pkg.GetEnv("REDIS_PASSWORD", "")
}

// ConnectRedis stabilisce una connessione con Redis
func ConnectRedis() (*redis.Client, error) {
	redisURL := RedisURL()
	redisPassword := RedisPassword()
	redisDB := 1

	// Crea il client Redis
	client := redis.NewClient(&redis.Options{
		Addr:         redisURL,
		Password:     redisPassword,
		DB:           redisDB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Testa la connessione
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("errore nel ping di Redis: %w", err)
	}

	log.Printf("Connessione a Redis stabilita con successo")

	return client, nil
}

// CloseRedis chiude la connessione a Redis
func CloseRedis(client *redis.Client) error {
	if client != nil {
		log.Println("Chiusura connessione Redis...")
		return client.Close()
	}
	return nil
}
