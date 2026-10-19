package http

import (
	"net/http"

	"encurtador-url/internal/http/handlers"
	"encurtador-url/internal/http/middleware"
	"encurtador-url/internal/service"
	"encurtador-url/pkg/logger"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(urlService *service.URLService, log *logger.Logger) http.Handler {
	r := chi.NewRouter()

	// Middlewares globais
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Rate limit simples
	r.Use(middleware.RateLimit(10)) // 10 req/min por IP

	//Página HTML
	r.Get("/", handlers.ServeIndex)

	// Handlers
	shortenHandler := handlers.NewShortenHandler(urlService, log)
	redirectHandler := handlers.NewRedirectHandler(urlService, log)

	// Rotas
	r.Post("/shorten", shortenHandler.Handle)
	r.Get("/{code}", redirectHandler.Handle)

	return r
}
