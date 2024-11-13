package models

import "time"

type Auth struct {
	UserID   string
	Password string
	Name     string
	Role     string
}

type Token struct {
	AccessToken string
	ExpiresAt   time.Time
}
