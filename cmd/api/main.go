package main

import (
	"log"
	net "net/http"

	"encurtador-url/internal/config"
	"encurtador-url/internal/http"
	"encurtador-url/internal/repository"
	"encurtador-url/internal/service"
	"encurtador-url/pkg/logger"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("erro ao carregar configurações: %v", err)
	}

	// Inicializar logger
	l := logger.New(cfg.Env)

	// Conectar ao banco SQLite
	db, err := repository.NewSQLiteDB(cfg.DatabasePath)
	if err != nil {
		l.Fatal("erro ao conectar ao banco", "error", err)
	}
	defer db.Close()

	// Inicializar repositório
	urlRepo := repository.NewURLRepository(db)

	// Inicializar serviço
	urlService := service.NewURLService(urlRepo)

	// Inicializar router
	router := http.NewRouter(urlService, l)

	l.Info("Servidor iniciado", "port", cfg.Port)

	// Subir servidor HTTP
	if err := net.ListenAndServe(":"+cfg.Port, router); err != nil {
		l.Fatal("erro ao iniciar servidor", "error", err)
	}
}
