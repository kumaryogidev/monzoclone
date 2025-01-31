package models

import (
	"time"

	"github.com/gocql/gocql"
)

type User struct {
	ID        gocql.UUID `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
}
