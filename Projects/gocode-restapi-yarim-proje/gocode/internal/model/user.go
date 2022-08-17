package model

import "time"

// type User struct {}

type (
	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
