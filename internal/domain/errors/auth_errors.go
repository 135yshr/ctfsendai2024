package errors

import "errors"

var (
	ErrInvalidPassword   = errors.New("パスワードが一致しません")
	ErrInvalidRequest    = errors.New("リクエストの形式が不正です")
	ErrUserNotFound      = errors.New("ユーザーが見つかりません")
	ErrNilAuth           = errors.New("認証情報がnilです")
	ErrInvalidToken      = errors.New("無効なトークンです")
	ErrTokenExpired      = errors.New("トークンの有効期限が切れています")
	ErrInvalidSignMethod = errors.New("無効な署名方式です")
	ErrInvalidUserID     = errors.New("無効なユーザーIDです")
)
