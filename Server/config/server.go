package config

import (
	"context"
	"log"
	"net/http"
	"oniplu/pkg"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ServerConfig contiene la configurazione del server HTTP
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// NewServerConfig crea una nuova configurazione del server con valori di default sicuri
func Config() *ServerConfig {
	return &ServerConfig{
		Port:         pkg.GetEnv("PORT", "3000"),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// StartServer avvia il server HTTP con graceful shutdown
// Equivalente alla funzione startServer di Express
func StartServer(handler http.Handler) {
	config := Config()

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: handler,
		// Configurazioni di sicurezza
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		// Limita la dimensione degli header per sicurezza
		ReadHeaderTimeout: 5 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	// Avvia il server in una goroutine
	go func() {
		log.Printf("Server avviato su porta %s", config.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Errore avvio server: %v", err)
		}
	}()

	// Gestisce il graceful shutdown
	gracefulShutdown(server)
}

// gracefulShutdown gestisce lo spegnimento sicuro del server
func gracefulShutdown(server *http.Server) {
	// Canale per catturare segnali del sistema operativo
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Blocca finché non riceve un segnale
	sig := <-quit
	log.Printf("Ricevuto segnale %v, iniziando graceful shutdown...", sig)

	// Crea un context con timeout per il graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Spegne il server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Errore durante graceful shutdown: %v", err)
		log.Printf("Forzo la chiusura del server...")
		if err := server.Close(); err != nil {
			log.Fatalf("Errore durante chiusura forzata: %v", err)
		}
	}

	log.Println("Server spento correttamente")
}
