package errors

import "errors"

var (
	ErrInvalidPassword     = errors.New("パスワードが一致しません")
	ErrInvalidRequest      = errors.New("リクエストの形式が不正です")
	ErrUserNotFound        = errors.New("ユーザーが見つかりません")
	ErrNilAuth             = errors.New("認証情報がnilです")
	ErrInvalidToken        = errors.New("無効なトークンです")
	ErrTokenExpired        = errors.New("トークンの有効期限が切れています")
	ErrInvalidSignMethod   = errors.New("無効な署名方式です")
	ErrInvalidUserID       = errors.New("無効なユーザーIDです")
	ErrInvalidRole         = errors.New("無効なロールです")
	ErrInvalidName         = errors.New("無効な名前です")
	ErrInvalidUser         = errors.New("無効なユーザーです")
	ErrForbidden           = errors.New("アクセスが禁止されています")
	ErrPlanNotFound        = errors.New("プランが見つかりません")
	ErrInvalidSecretAnswer = errors.New("秘密の質問の回答が不正です")
)
