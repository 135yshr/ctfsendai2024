package models

import "time"

type Auth struct {
	UserID   string
	Username string
	Password string
	Role     string
}

type Token struct {
	AccessToken string
	ExpiresAt   time.Time
}
