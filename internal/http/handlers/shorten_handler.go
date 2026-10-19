package handlers

import (
	"encoding/json"
	"net/http"

	"encurtador-url/internal/service"
	"encurtador-url/pkg/logger"
)

type ShortenHandler struct {
	service *service.URLService
	log     *logger.Logger
}

func NewShortenHandler(s *service.URLService, l *logger.Logger) *ShortenHandler {
	return &ShortenHandler{
		service: s,
		log:     l,
	}
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortCode string `json:"short_code"`
}

func (h *ShortenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest

	// Decodificar JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("erro ao decodificar JSON", "error", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validar campo
	if req.URL == "" {
		http.Error(w, "campo 'url' é obrigatório", http.StatusBadRequest)
		return
	}

	// Chamar serviço
	url, err := h.service.ShortenURL(req.URL)
	if err != nil {
		h.log.Error("erro ao encurtar URL", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Resposta
	resp := shortenResponse{
		ShortCode: url.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	h.log.Info("URL encurtada com sucesso", "short_code", url.ShortCode)
}
