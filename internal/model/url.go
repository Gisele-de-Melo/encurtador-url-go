package model

import "time"

type URL struct {
	ID        int64     `json:"id"`
	Original  string    `json:"original"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
}
