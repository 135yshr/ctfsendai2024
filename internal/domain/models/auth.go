package models

import "time"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Auth struct {
	UserID   string
	Password string
	Name     string
	Role     Role
}

type Token struct {
	AccessToken string
	ExpiresAt   time.Time
}
