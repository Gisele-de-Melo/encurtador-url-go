package repository

import (
	"database/sql"
	"errors"

	"encurtador-url/internal/model"

	_ "modernc.org/sqlite"
)

type URLRepository struct {
	db *sql.DB
}

func NewSQLiteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// Criar tabela se não existir
	schema := `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original TEXT NOT NULL,
        short_code TEXT NOT NULL UNIQUE,
        created_at DATETIME NOT NULL
    );
    `

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url *model.URL) error {
	query := `
        INSERT INTO urls (original, short_code, created_at)
        VALUES (?, ?, ?)
    `

	_, err := r.db.Exec(query, url.Original, url.ShortCode, url.CreatedAt)
	return err
}

func (r *URLRepository) FindByShortCode(code string) (*model.URL, error) {
	query := `
        SELECT id, original, short_code, created_at
        FROM urls
        WHERE short_code = ?
    `

	row := r.db.QueryRow(query, code)

	var u model.URL
	err := row.Scan(&u.ID, &u.Original, &u.ShortCode, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // não encontrado
		}
		return nil, err
	}

	return &u, nil
}
