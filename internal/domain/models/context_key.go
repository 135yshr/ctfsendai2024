package models

// ContextKey はコンテキストのキーを表す型です.
type ContextKey string

const (
	// UserContextKey はユーザー情報を保存するためのコンテキストキーです.
	UserContextKey ContextKey = "user"
)
