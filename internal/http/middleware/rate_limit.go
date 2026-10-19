package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type clientData struct {
	requests int
	resetAt  time.Time
}

var (
	clients = make(map[string]*clientData)
	mu      sync.Mutex
)

// RateLimit limita requisições por IP por minuto
func RateLimit(maxRequests int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ip := getIP(r)

			mu.Lock()
			data, exists := clients[ip]

			// Se não existe, cria
			if !exists {
				data = &clientData{
					requests: 1,
					resetAt:  time.Now().Add(time.Minute),
				}
				clients[ip] = data
				mu.Unlock()
				next.ServeHTTP(w, r)
				return
			}

			// Se o minuto já passou, reseta
			if time.Now().After(data.resetAt) {
				data.requests = 1
				data.resetAt = time.Now().Add(time.Minute)
				mu.Unlock()
				next.ServeHTTP(w, r)
				return
			}

			// Se excedeu o limite
			if data.requests >= maxRequests {
				mu.Unlock()
				http.Error(w, "limite de requisições excedido", http.StatusTooManyRequests)
				return
			}

			// Incrementa e segue
			data.requests++
			mu.Unlock()

			next.ServeHTTP(w, r)
		})
	}
}

// Captura o IP real do cliente
func getIP(r *http.Request) string {
	// Se estiver atrás de proxy
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return forwarded
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
