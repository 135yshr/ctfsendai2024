package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type userDatabase struct {
	Users []struct {
		ID             string `json:"id"`
		Password       string `json:"password"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		Phone          string `json:"phone"`
		Role           string `json:"role"`
		SecretQuestion string `json:"secret_question"`
		SecretAnswer   string `json:"secret_answer"`
	} `json:"users"`
}

type UserRepository struct {
	users map[string]*models.User
}

func NewUserRepository(dbPath string) (repositories.UserRepository, error) {
	repo := &UserRepository{users: make(map[string]*models.User)}
	if err := repo.loadUsers(dbPath); err != nil {
		return nil, fmt.Errorf("ユーザー情報の読み込みに失敗しました: %w", err)
	}

	return repo, nil
}

func (r *UserRepository) loadUsers(dbPath string) error {
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました: %w", err)
	}

	var database userDatabase
	if err = json.Unmarshal(data, &database); err != nil {
		return fmt.Errorf("設定ファイルのパースに失敗しました: %w", err)
	}

	for _, user := range database.Users {
		r.users[user.ID] = &models.User{
			ID:             user.ID,
			Password:       user.Password,
			Name:           user.Name,
			Email:          user.Email,
			Phone:          user.Phone,
			Role:           models.Role(user.Role),
			SecretQuestion: user.SecretQuestion,
			SecretAnswer:   user.SecretAnswer,
		}
	}

	return nil
}

// FindByUserID はユーザー名からユーザーを検索します.
func (r *UserRepository) FindByUserID(_ context.Context, userID string) (*models.User, error) {
	if user, exists := r.users[userID]; exists {
		return user, nil
	}

	return nil, errors.ErrUserNotFound
}

// Store はユーザーを保存します.
func (r *UserRepository) Store(_ context.Context, user *models.User) error {
	if user == nil {
		return errors.ErrInvalidUser
	}
	r.users[user.ID] = user

	return nil
}
