package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
	"time"

	"encurtador-url/internal/model"
	"encurtador-url/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

// ShortenURL cria um shortcode, valida a URL e salva no banco
func (s *URLService) ShortenURL(original string) (*model.URL, error) {
	// Validar URL
	if !isValidURL(original) {
		return nil, errors.New("URL inválida")
	}

	// Gerar shortcode único
	shortCode, err := s.generateShortCode()
	if err != nil {
		return nil, err
	}

	u := &model.URL{
		Original:  original,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
	}

	// Salvar no banco
	err = s.repo.Save(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetOriginalURL retorna a URL original pelo shortcode
func (s *URLService) GetOriginalURL(code string) (*model.URL, error) {
	return s.repo.FindByShortCode(code)
}

// ----------------------
// Funções auxiliares
// ----------------------

func isValidURL(raw string) bool {
	if raw == "" {
		return false
	}

	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func (s *URLService) generateShortCode() (string, error) {
	// 4 bytes → ~6 caracteres base64
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	code := base64.URLEncoding.EncodeToString(b)
	code = strings.TrimRight(code, "=") // remover '=' do base64

	return code, nil
}
