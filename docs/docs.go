// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/api/folder": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Folder"
                ],
                "summary": "Получить все папки пользователя",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/domain.Folder"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
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
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Folder"
                ],
                "summary": "Создать папку",
                "parameters": [
                    {
                        "description": "folder",
                        "name": "folder",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.FolderCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/folder/spreading": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Folder"
                ],
                "summary": "Добавить раскидку в папку",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.FolderAddSpreadingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/folder/strategy": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Folder"
                ],
                "summary": "Добавить стратегию в папку",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.FolderAddStrategyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/payment/create/tinkoff": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Создать платеж",
                "parameters": [
                    {
                        "description": "paymentInfo",
                        "name": "paymentInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PaymentCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.PaymentCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/profile": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Получить информацию о пользователе",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Profile"
                        }
                    }
                }
            }
        },
        "/api/spreading": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spreading"
                ],
                "summary": "Получить несколько раскидок",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ids",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/domain.Spreading"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
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
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spreading"
                ],
                "summary": "Создать раскидку",
                "parameters": [
                    {
                        "description": "spreading",
                        "name": "spreading",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SpreadingCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Spreading"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/spreading/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spreading"
                ],
                "summary": "Получить одну раскидку по айди",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Spreading"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spreading"
                ],
                "summary": "Обновить раскидку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SpreadingUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/strategy": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "Получить несколько стратегий",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ids",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/domain.Strategy"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
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
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "Создать стратегию",
                "parameters": [
                    {
                        "description": "strategy",
                        "name": "strategy",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.StrategyCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.StrategyCreateRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/strategy/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "Получить одну стратегию по айди",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Strategy"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "Обновить стратегию",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.StrategyUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Вход по почте и паролю",
                "parameters": [
                    {
                        "description": "login request",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Refresh"
                ],
                "summary": "Обновить токены",
                "parameters": [
                    {
                        "description": "refreshToken",
                        "name": "refreshToken",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reset/password": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ResetPassword"
                ],
                "summary": "Смена пароля",
                "parameters": [
                    {
                        "description": "token and name request",
                        "name": "tokenRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ResetPassword"
                ],
                "summary": "Отправить запрос на смену пароля",
                "parameters": [
                    {
                        "description": "create code request",
                        "name": "createResetPasswordRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ResetPassowrdCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reset/password/confirm": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ResetPassword"
                ],
                "summary": "Подтверждение почты по коду",
                "parameters": [
                    {
                        "description": "code request",
                        "name": "codeRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ResetPasswordConfirmRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ResetPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup/confirm": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Signup"
                ],
                "summary": "Подтверждение почты",
                "parameters": [
                    {
                        "description": "code request",
                        "name": "codeRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ConfirmRegisterCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Signup"
                ],
                "summary": "Регистрация по почте и паролю",
                "parameters": [
                    {
                        "description": "sign up request",
                        "name": "signUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ConfirmRegisterCodeRequest": {
            "type": "object",
            "required": [
                "code",
                "email",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 3
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 2
                }
            }
        },
        "domain.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.Folder": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100
                },
                "spreadouts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "strategies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "domain.FolderAddSpreadingRequest": {
            "type": "object",
            "required": [
                "folder_id",
                "spreading_id"
            ],
            "properties": {
                "folder_id": {
                    "type": "string"
                },
                "spreading_id": {
                    "type": "string"
                }
            }
        },
        "domain.FolderAddStrategyRequest": {
            "type": "object",
            "required": [
                "folder_id",
                "strategy_id"
            ],
            "properties": {
                "folder_id": {
                    "type": "string"
                },
                "strategy_id": {
                    "type": "string"
                }
            }
        },
        "domain.FolderCreateRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "domain.PaymentCreateRequest": {
            "type": "object",
            "required": [
                "days",
                "email"
            ],
            "properties": {
                "days": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 3
                }
            }
        },
        "domain.PaymentCreated": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "domain.Profile": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pro": {
                    "$ref": "#/definitions/domain.UserPro"
                }
            }
        },
        "domain.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "domain.RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "domain.ResetPassowrdCreate": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "domain.ResetPasswordConfirmRequest": {
            "type": "object",
            "required": [
                "code",
                "email"
            ],
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "domain.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "password",
                "token"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.ResetPasswordResponse": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                }
            }
        },
        "domain.Spreading": {
            "type": "object",
            "required": [
                "elements",
                "map_name",
                "name"
            ],
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "id": {
                    "type": "string"
                },
                "map_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "domain.SpreadingCreateRequest": {
            "type": "object",
            "required": [
                "elements",
                "map_name",
                "name"
            ],
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "map_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "domain.SpreadingUpdateRequest": {
            "type": "object",
            "required": [
                "elements",
                "map_name"
            ],
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "map_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "domain.Strategy": {
            "type": "object",
            "required": [
                "map_name",
                "name",
                "parts"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "map_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "parts": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "domain.StrategyCreateRequest": {
            "type": "object",
            "required": [
                "map_name",
                "name",
                "parts"
            ],
            "properties": {
                "map_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "parts": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "domain.StrategyUpdateRequest": {
            "type": "object",
            "required": [
                "map_name",
                "parts"
            ],
            "properties": {
                "map_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "parts": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "domain.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.UserPro": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "until": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Ston Tactics",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
