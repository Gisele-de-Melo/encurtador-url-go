package handlers

import (
	"net/http"

	"encurtador-url/internal/service"
	"encurtador-url/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type RedirectHandler struct {
	service *service.URLService
	log     *logger.Logger
}

func NewRedirectHandler(s *service.URLService, l *logger.Logger) *RedirectHandler {
	return &RedirectHandler{
		service: s,
		log:     l,
	}
}

func (h *RedirectHandler) Handle(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	if code == "" {
		http.Error(w, "shortcode não informado", http.StatusBadRequest)
		return
	}

	// Buscar URL original
	url, err := h.service.GetOriginalURL(code)
	if err != nil {
		h.log.Error("erro ao buscar URL", "error", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}

	if url == nil {
		http.Error(w, "URL não encontrada", http.StatusNotFound)
		return
	}

	// Redirecionar
	http.Redirect(w, r, url.Original, http.StatusFound)

	h.log.Info("redirecionamento realizado", "short_code", code, "to", url.Original)
}
