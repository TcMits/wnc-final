// Package partner GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package partner

import "github.com/swaggo/swag"

const docTemplatepartner = `{
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
        "/auth": {
            "post": {
                "description": "Authenticate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Authenticate",
                "operationId": "authenticate",
                "parameters": [
                    {
                        "description": "Authenticate",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/partner.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/partner.tokenPairResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    }
                }
            }
        },
        "/bank-accounts": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get bank account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank account"
                ],
                "summary": "Get bank account",
                "operationId": "bankaccount-listing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account number of bank account",
                        "name": "account_number",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/partner.bankAccountResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    }
                }
            }
        },
        "/options": {
            "get": {
                "description": "Show all options",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Option"
                ],
                "summary": "Show options",
                "operationId": "option",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/partner.optionsResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Create a transaction",
                "operationId": "transaction-create",
                "parameters": [
                    {
                        "description": "Create a transaction",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/partner.transactionCreateReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/partner.transactionResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/validate": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Validate before create transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Validate before create transaction",
                "operationId": "transaction-validate",
                "parameters": [
                    {
                        "description": "Validate before create transaction",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/partner.transactionCreateReq"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/partner.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "partner.bankAccountResp": {
            "type": "object",
            "properties": {
                "account_number": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "partner.errorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "partner.loginRequest": {
            "type": "object",
            "required": [
                "api_key"
            ],
            "properties": {
                "api_key": {
                    "type": "string"
                }
            }
        },
        "partner.optionsResp": {
            "type": "object",
            "properties": {
                "actor_types": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "partner.tokenPairResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "partner.transactionCreateReq": {
            "type": "object",
            "required": [
                "amount",
                "description",
                "fee_paid_by",
                "sender_bank_account_number",
                "sender_name",
                "signature",
                "token"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "fee_paid_by": {
                    "type": "string"
                },
                "sender_bank_account_number": {
                    "type": "string"
                },
                "sender_name": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "partner.transactionResp": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "create_time": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "receiver_bank_account_number": {
                    "type": "string"
                },
                "receiver_bank_name": {
                    "type": "string"
                },
                "receiver_id": {
                    "type": "string"
                },
                "receiver_name": {
                    "type": "string"
                },
                "sender_bank_account_number": {
                    "type": "string"
                },
                "sender_bank_name": {
                    "type": "string"
                },
                "sender_id": {
                    "type": "string"
                },
                "sender_name": {
                    "type": "string"
                },
                "source_transaction_id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/transaction.Status"
                },
                "transaction_type": {
                    "$ref": "#/definitions/transaction.TransactionType"
                },
                "update_time": {
                    "type": "string"
                }
            }
        },
        "transaction.Status": {
            "type": "string",
            "enum": [
                "draft",
                "draft",
                "verified",
                "success"
            ],
            "x-enum-varnames": [
                "DefaultStatus",
                "StatusDraft",
                "StatusVerified",
                "StatusSuccess"
            ]
        },
        "transaction.TransactionType": {
            "type": "string",
            "enum": [
                "internal",
                "external"
            ],
            "x-enum-varnames": [
                "TransactionTypeInternal",
                "TransactionTypeExternal"
            ]
        }
    }
}`

// SwaggerInfopartner holds exported Swagger Info so clients can modify it
var SwaggerInfopartner = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "partner",
	SwaggerTemplate:  docTemplatepartner,
}

func init() {
	swag.Register(SwaggerInfopartner.InstanceName(), SwaggerInfopartner)
}
