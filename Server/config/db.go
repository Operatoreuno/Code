package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"oniplu/pkg"
	"time"

	_ "github.com/lib/pq"
)

// DatabaseURL restituisce l'URL del database dalle variabili d'ambiente
func DatabaseURL() string {
	return pkg.GetEnv("DB_URL", "postgres://dmnc@localhost:5432/db?sslmode=disable")
}

// ConnectDatabase stabilisce una connessione con il database PostgreSQL
func ConnectDatabase() (*sql.DB, error) {
	dbURL := DatabaseURL()

	// Apri la connessione al database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("errore nell'apertura del database: %w", err)
	}

	// Configura il pool di connessioni
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Testa la connessione
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("errore nel ping del database: %w", err)
	}

	log.Printf("Connessione al database PostgreSQL stabilita con successo")

	return db, nil
}

// CloseDatabase chiude la connessione al database
func CloseDatabase(db *sql.DB) error {
	if db != nil {
		log.Println("Chiusura connessione database...")
		return db.Close()
	}
	return nil
}
