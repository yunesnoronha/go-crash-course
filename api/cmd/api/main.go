package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yunesnoronha/go-crash-course/api/internal/config"
	"github.com/yunesnoronha/go-crash-course/api/internal/handlers"
	"github.com/yunesnoronha/go-crash-course/api/internal/middleware"
)

func main() {
	// Carregar configuração
	cfg := config.Load()

	// Criar router
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", handlers.HealthCheck)

	// API routes
	mux.HandleFunc("/api/v1/users", handlers.UsersHandler)
	mux.HandleFunc("/api/v1/users/", handlers.UserHandler) // Para rotas com ID

	// Aplicar middlewares
	handler := middleware.Logger(
		middleware.Recovery(
			middleware.CORS(mux),
		),
	)

	// Configurar servidor
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Canal para capturar sinais do sistema
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Iniciar servidor em uma goroutine
	go func() {
		log.Printf("🚀 Servidor iniciado na porta %s", cfg.Port)
		log.Printf("📍 URL: http://localhost:%s", cfg.Port)
		log.Printf("🏥 Health: http://localhost:%s/health", cfg.Port)
		log.Printf("📚 API: http://localhost:%s/api/v1", cfg.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Aguardar sinal de interrupção
	<-stop

	// Graceful shutdown
	log.Println("\n🛑 Encerrando servidor gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao encerrar servidor: %v", err)
	}

	log.Println("✅ Servidor encerrado com sucesso")
}
