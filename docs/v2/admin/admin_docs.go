// Package admin GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package admin

import "github.com/swaggo/swag"

const docTemplateadmin = `{
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
        "/api/admin/v1/employees": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Show employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Show employee",
                "operationId": "employee-listing",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.EntitiesResponseTemplate-admin_employeeResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
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
                "description": "Create a employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Create a employee",
                "operationId": "employee-create",
                "parameters": [
                    {
                        "description": "Create a employee",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.employeeCreateReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/admin.employeeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/v1/employees/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Get a employee",
                "operationId": "employee-get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.employeeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
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
                "description": "Update a employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Update a employee",
                "operationId": "employee-update",
                "parameters": [
                    {
                        "description": "Update a employee",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.employeeUpdateReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID of employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.employeeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Delete a employee",
                "operationId": "employee-delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/v1/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.tokenPairResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Logout",
                "operationId": "logout",
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/v1/me/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Me"
                ],
                "summary": "Get profile",
                "operationId": "me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.meResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/v1/token": {
            "post": {
                "description": "Renew token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Renew token",
                "operationId": "renewtoken",
                "parameters": [
                    {
                        "description": "Renew token",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.renewTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.tokenPairResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/v1/transactions": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Show transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Show transactions",
                "operationId": "transaction-listing",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "True if sort ascent by update_time otherwise ignored",
                        "name": "update_time",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "True if sort descent by update_time otherwise ignored",
                        "name": "-update_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Date start",
                        "name": "date_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Date end",
                        "name": "date_end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bank name",
                        "name": "bank_name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.EntitiesResponseTemplate-admin_transactionResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/v1/transactions/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get a transaction",
                "operationId": "transaction-get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of transaction",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.transactionResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/admin.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.EntitiesResponseTemplate-admin_employeeResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "next": {
                    "type": "string"
                },
                "previous": {
                    "type": "string"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.employeeResponse"
                    }
                }
            }
        },
        "admin.EntitiesResponseTemplate-admin_transactionResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "next": {
                    "type": "string"
                },
                "previous": {
                    "type": "string"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.transactionResp"
                    }
                }
            }
        },
        "admin.employeeCreateReq": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "username"
            ],
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.employeeResponse": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.employeeUpdateReq": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.errorResponse": {
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
        "admin.loginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.meResponse": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.renewTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "admin.tokenPairResponse": {
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
        "admin.transactionResp": {
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

// SwaggerInfoadmin holds exported Swagger Info so clients can modify it
var SwaggerInfoadmin = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "admin",
	SwaggerTemplate:  docTemplateadmin,
}

func init() {
	swag.Register(SwaggerInfoadmin.InstanceName(), SwaggerInfoadmin)
}
