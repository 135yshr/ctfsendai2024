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
                            "$ref": "#/definitions/dto.LoginRequest"
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
                            "$ref": "#/definitions/presenters.PresentError"
                        }
                    },
                    "401": {
                        "description": "パスワードが一致しない",
                        "schema": {
                            "$ref": "#/definitions/presenters.PresentError"
                        }
                    },
                    "500": {
                        "description": "サーバー内部エラー",
                        "schema": {
                            "$ref": "#/definitions/presenters.PresentError"
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
                "description": "検索条件に基づいてプラン一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "プラン一覧取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "開始日 (YYYY-MM-DD)",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "終了日 (YYYY-MM-DD)",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ステータス",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenters.PlansResponse"
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
        }
    },
    "definitions": {
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "user_id"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "integer"
                }
            }
        },
        "dto.PlanResponse": {
            "description": "プラン情報の詳細.",
            "type": "object",
            "properties": {
                "description": {
                    "description": "プランの説明\n@Example \"基本的なサービスが含まれるプランです\"",
                    "type": "string",
                    "example": "基本的なサービスが含まれるプランです"
                },
                "duration": {
                    "description": "プランの期間（日数）\n@Example 30",
                    "type": "integer",
                    "example": 30
                },
                "id": {
                    "description": "プランID\n@Example \"plan123\"",
                    "type": "string",
                    "example": "plan123"
                },
                "name": {
                    "description": "プラン名\n@Example \"スタンダードプラン\"",
                    "type": "string",
                    "example": "スタンダードプラン"
                },
                "price": {
                    "description": "プランの価格\n@Example 1000",
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
                    "description": "予約終了時間\n@Example \"2024-03-20T11:00:00Z\"",
                    "type": "string",
                    "example": "2024-03-20T11:00:00Z"
                },
                "id": {
                    "description": "予約ID\n@Example \"rsv_123456\"",
                    "type": "string",
                    "example": "rsv_123456"
                },
                "plan": {
                    "description": "プラン情報\n@Example {\"id\": \"plan123\", \"name\": \"スタンダードプラン\"}",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.PlanResponse"
                        }
                    ]
                },
                "start_time": {
                    "description": "予約開始時間\n@Example \"2024-03-20T10:00:00Z\"",
                    "type": "string",
                    "example": "2024-03-20T10:00:00Z"
                },
                "status": {
                    "description": "予約ステータス\n@Example \"confirmed\"",
                    "type": "string",
                    "example": "confirmed"
                },
                "user": {
                    "description": "ユーザー情報\n@Example {\"id\": \"user123\", \"name\": \"山田太郎\"}",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    ]
                },
                "user_id": {
                    "description": "ユーザーID\n@Example \"user123\"",
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
                    "description": "メールアドレス\n@Example \"taro.yamada@example.com\"",
                    "type": "string",
                    "example": "taro.yamada@example.com"
                },
                "id": {
                    "description": "ユーザーID\n@Example \"user123\"",
                    "type": "string",
                    "example": "user123"
                },
                "name": {
                    "description": "ユーザー名\n@Example \"山田太郎\"",
                    "type": "string",
                    "example": "山田太郎"
                },
                "phone": {
                    "description": "電話番号\n@Example \"090-1234-5678\"",
                    "type": "string",
                    "example": "090-1234-5678"
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
                    "description": "ステータス\n@Example \"success\"",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.PlansResponse": {
            "description": "プラン一覧のレスポンス.",
            "type": "object",
            "properties": {
                "data": {
                    "description": "プランデータ",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PlanResponse"
                    }
                },
                "status": {
                    "description": "ステータス\n@Example \"success\"",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.PresentError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "presenters.ReservationResponse": {
            "description": "予約のレスポンス.",
            "type": "object",
            "properties": {
                "data": {
                    "description": "予約データ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.ReservationResponse"
                        }
                    ]
                },
                "status": {
                    "description": "ステータス\n@Example \"success\"",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenters.ReservationsResponse": {
            "description": "予約一覧のレスポンス.",
            "type": "object",
            "properties": {
                "data": {
                    "description": "予約データ",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ReservationResponse"
                    }
                },
                "status": {
                    "description": "ステータス\n@Example \"success\"",
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
                    "description": "エラーメッセージ\n@Example \"プランが見つかりません\"",
                    "type": "string",
                    "example": "プランが見つかりません"
                },
                "status": {
                    "description": "ステータス\n@Example \"error\"",
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "validators.CreateReservationRequest": {
            "type": "object",
            "required": [
                "plan_id",
                "start_date",
                "user_id"
            ],
            "properties": {
                "plan_id": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
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
	Host:             "localhost:8080",
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
