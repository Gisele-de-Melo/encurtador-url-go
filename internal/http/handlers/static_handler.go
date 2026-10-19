package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	// Caminho absoluto até o arquivo index.html
	path, err := filepath.Abs("public/index.html")
	if err != nil {
		http.Error(w, "Erro ao localizar arquivo HTML", http.StatusInternalServerError)
		return
	}

	// Verifica se o arquivo existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "Arquivo index.html não encontrado", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path)
}
