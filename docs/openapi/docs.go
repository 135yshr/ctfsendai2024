// Package openapi Code generated by swaggo/swag. DO NOT EDIT
package openapi

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "ユーザー名とパスワードを使用してログイン認証を行います",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "ユーザーログイン",
                "parameters": [
                    {
                        "description": "ログイン情報",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ログイン成功時のレスポンス",
                        "schema": {
                            "$ref": "#/definitions/presenters.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストの形式が不正",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "パスワードが一致しない",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "サーバー内部エラー",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/plans": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "ユーザーの権限に基づいて、指定された検索条件に一致するプラン一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "プラン一覧取得API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "開始日 (YYYY-MM-DD形式)",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "終了日 (YYYY-MM-DD形式)",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "プランのステータス (reserved/canceled/confirmed/pending)",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "プラン一覧の取得に成功",
                        "schema": {
                            "$ref": "#/definitions/presenters.PlansResponse"
                        }
                    },
                    "400": {
                        "description": "不正なリクエストパラメータ",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "認証エラー",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "サーバー内部エラー",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reservations": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "指定されたユーザーIDに紐づく予約の一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "ユーザーの予約一覧を取得",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 3,
                        "type": "string",
                        "description": "ユーザーID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenters.ReservationsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "ユーザーの新しい予約を作成します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "新しい予約を作成",
                "parameters": [
                    {
                        "description": "予約情報",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.CreateReservationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/presenters.ReservationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/secret-login": {
            "post": {
                "description": "秘密の質問の回答を使用してログイン認証を行います",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "秘密の質問によるログイン",
                "parameters": [
                    {
                        "description": "秘密の質問の回答情報",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.SecretLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ログイン成功時のレスポンス",
                        "schema": {
                            "$ref": "#/definitions/presenters.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストの形式が不正",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "秘密の質問の回答が一致しない",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "サーバー内部エラー",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/secret-question": {
            "get": {
                "description": "ユーザーIDに対応する秘密の質問を取得します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "秘密の質問の取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ユーザーID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "秘密の質問",
                        "schema": {
                            "$ref": "#/definitions/presenters.SecretQuestionResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストの形式が不正",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "ユーザーが見つからない",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "サーバー内部エラー",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "全ユーザーの情報を取得します（管理者のみ）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "ユーザー一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/presenters.UserResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "ログイン中のユーザー情報を取得します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "ログインユーザー情報取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenters.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.LoginResponse": {
            "description": "ログイン処理のレスポンス.",
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "アクセストークン",
                    "type": "string"
                },
                "expires_at": {
                    "description": "トークンの有効期限",
                    "type": "integer"
                }
            }
        },
        "dto.PlanResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "プランの説明",
                    "type": "string",
                    "example": "基本的なサービスが含まれるプランです"
                },
                "duration": {
                    "description": "プランの期間（日数）",
                    "type": "integer",
                    "example": 30
                },
                "id": {
                    "description": "プランID",
                    "type": "string",
                    "example": "plan123"
                },
                "name": {
                    "description": "プラン名",
                    "type": "string",
                    "example": "スタンダードプラン"
                },
                "price": {
                    "description": "プランの価格",
                    "type": "integer",
                    "example": 1000
                }
            }
        },
        "dto.ReservationResponse": {
            "description": "予約情報の詳細.",
            "type": "object",
            "properties": {
                "end_time": {
                    "description": "予約終了時間",
                    "type": "string",
                    "example": "2024-03-20T11:00:00Z"
                },
                "id": {
                    "description": "予約ID",
                    "type": "string",
                    "example": "rsv_123456"
                },
                "plan": {
                    "description": "プラン情報",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.PlanResponse"
                        }
                    ]
                },
                "start_time": {
                    "description": "予約開始時間",
                    "type": "string",
                    "example": "2024-03-20T10:00:00Z"
                },
                "status": {
                    "description": "予約ステータス",
                    "type": "string",
                    "example": "confirmed"
                },
                "user": {
                    "description": "ユーザー情報",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    ]
                },
                "user_id": {
                    "description": "ユーザーID",
                    "type": "string",
                    "example": "user123"
                }
            }
        },
        "dto.UserResponse": {
            "description": "ユーザー情報の詳細.",
            "type": "object",
            "properties": {
                "email": {
                    "description": "メールアドレス",
                    "type": "string",
                    "example": "taro.yamada@example.com"
                },
                "id": {
                    "description": "ユーザーID",
                    "type": "string",
                    "example": "user123"
                },
                "name": {
                    "description": "ユーザー名",
                    "type": "string",
                    "example": "山田太郎"
                },
                "phone": {
                    "description": "電話番号",
                    "type": "string",
                    "example": "FLAG_dSQVRVTEFUSU9OU19GT1JfRklOSVNISU5H"
                }
            }
        },
        "presenters.LoginResponse": {
            "description": "ログイン処理のレスポンス.",
            "type": "object",
            "properties": {
                "data": {
                    "description": "ログインデータ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    ]
                },
                "status": {
                    "description": "ステータス",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.PlansResponse": {
            "description": "プラン一覧のレスポンス形式を定義する.",
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data プラン情報一覧",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PlanResponse"
                    }
                },
                "status": {
                    "description": "Status レスポンスのステータス\n\t@Example\t\"success\"",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.ReservationResponse": {
            "description": "予約のレスポンス",
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data 予約データ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.ReservationResponse"
                        }
                    ]
                },
                "status": {
                    "description": "Status レスポンスのステータス",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.ReservationsResponse": {
            "description": "予約一覧のレスポンス",
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data 予約データ一覧",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ReservationResponse"
                    }
                },
                "status": {
                    "description": "Status レスポンスのステータス",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.SecretQuestionResponse": {
            "description": "秘密の質問のレスポンス.",
            "type": "object",
            "properties": {
                "secret_question": {
                    "description": "秘密の質問",
                    "type": "string"
                }
            }
        },
        "presenters.UserResponse": {
            "description": "ユーザー情報の単一レスポンス形式を定義する.",
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data ユーザー情報",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    ]
                },
                "status": {
                    "description": "Status レスポンスのステータス",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "response.ErrorResponse": {
            "description": "エラー情報のレスポンス.",
            "type": "object",
            "properties": {
                "error": {
                    "description": "エラーメッセージ",
                    "type": "string",
                    "example": "プランが見つかりません"
                },
                "status": {
                    "description": "ステータス",
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "validators.CreateReservationRequest": {
            "description": "新規予約を作成するためのリクエストパラメータ.",
            "type": "object",
            "required": [
                "plan_id",
                "start_date",
                "user_id"
            ],
            "properties": {
                "plan_id": {
                    "description": "プランID",
                    "type": "string"
                },
                "start_date": {
                    "description": "予約開始日時",
                    "type": "string"
                },
                "user_id": {
                    "description": "ユーザーID",
                    "type": "string"
                }
            }
        },
        "validators.LoginRequest": {
            "description": "ユーザーIDとパスワードによるログインリクエスト.",
            "type": "object",
            "required": [
                "password",
                "user_id"
            ],
            "properties": {
                "password": {
                    "description": "パスワード",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "user_id": {
                    "description": "ユーザーID",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3,
                    "example": "u00100"
                }
            }
        },
        "validators.SecretLoginRequest": {
            "description": "秘密の質問の回答によるログインリクエスト.",
            "type": "object",
            "required": [
                "secret_answer",
                "user_id"
            ],
            "properties": {
                "secret_answer": {
                    "description": "秘密の質問の回答",
                    "type": "string",
                    "minLength": 1
                },
                "user_id": {
                    "description": "ユーザーID",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Bearer Tokenによる認証",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "予約管理システム API",
	Description:      "予約管理システムのRESTful API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
