package models

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	Role           Role   `json:"role"`
	SecretQuestion string `json:"secret_question"`
	SecretAnswer   string `json:"secret_answer"`
}
