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
        "/daftar": {
            "post": {
                "description": "Mendaftarkan nasabah baru dengan nama, NIK, no_hp, dan tipe akun",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Register new nasabah",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/saldo/{no_rekening}": {
            "get": {
                "description": "Melihat saldo nasabah berdasarkan nomor rekening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get account balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nomor Rekening",
                        "name": "no_rekening",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SaldoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.SaldoResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.SaldoResponse"
                        }
                    }
                }
            }
        },
        "/tabung": {
            "post": {
                "description": "Menyetor saldo ke akun berdasarkan nomor rekening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Deposit saldo ke akun nasabah",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DepositRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.DepositResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tarik": {
            "post": {
                "description": "Menarik dana dari rekening tabungan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Withdraw money from an account",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.WithdrawRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.WithdrawResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.WithdrawResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.WithdrawResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.DepositRequest": {
            "type": "object",
            "required": [
                "no_rekening",
                "nominal"
            ],
            "properties": {
                "no_rekening": {
                    "type": "string"
                },
                "nominal": {
                    "type": "number"
                }
            }
        },
        "dto.DepositResponse": {
            "type": "object",
            "properties": {
                "saldo": {
                    "type": "number"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "remark": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterRequest": {
            "type": "object",
            "required": [
                "account_type",
                "nama",
                "nik",
                "no_hp"
            ],
            "properties": {
                "account_type": {
                    "type": "string",
                    "enum": [
                        "savings",
                        "giro"
                    ]
                },
                "nama": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterResponse": {
            "type": "object",
            "properties": {
                "no_rekening": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                }
            }
        },
        "dto.SaldoResponse": {
            "type": "object",
            "properties": {
                "remark": {
                    "type": "string"
                },
                "saldo": {
                    "type": "number"
                }
            }
        },
        "dto.WithdrawRequest": {
            "type": "object",
            "required": [
                "no_rekening",
                "nominal"
            ],
            "properties": {
                "no_rekening": {
                    "type": "string"
                },
                "nominal": {
                    "type": "number"
                }
            }
        },
        "dto.WithdrawResponse": {
            "type": "object",
            "properties": {
                "remark": {
                    "type": "string"
                },
                "saldo": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Go Backend Test",
	Description:      "Test for managing accounts & transactions service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
