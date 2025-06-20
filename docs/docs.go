// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Maaw",
            "email": "maaw@mail.ru"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/gsm": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an entry  in the GSM table with the specified GUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gsm"
                ],
                "summary": "Update an entry",
                "parameters": [
                    {
                        "description": "GSM data",
                        "name": "GsmEntry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GsmeEntryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.IdEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add an entry to the GSM table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gsm"
                ],
                "summary": "Add an entry",
                "parameters": [
                    {
                        "description": "GSM data",
                        "name": "GsmEntry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GsmeEntryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.IdEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/gsm/date/{date}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Receive an entry with  a specified date from the GSM table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gsm"
                ],
                "summary": "Receive an entry",
                "parameters": [
                    {
                        "type": "string",
                        "format": "date",
                        "description": "Date in the format YYYY-MM-DD",
                        "name": "date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.GsmEntryResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/gsm/id/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Receive an entry with a specified ID from the GSM table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gsm"
                ],
                "summary": "Receive an entry",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GsmEntryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an entry with a specified ID from the GSM table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gsm"
                ],
                "summary": "Delete an entry",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.IdEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.IdEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "User identification and authentication. If successful, it returns an access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User identification and authentication",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.AccessToken"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/{path}": {
            "get": {
                "description": "Supports GET/POST/PUT/DELETE for any URL",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "A custom path handler is executed if no route matches",
                "parameters": [
                    {
                        "type": "string",
                        "description": "An arbitrary path",
                        "name": "path",
                        "in": "path"
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "Supports GET/POST/PUT/DELETE for any URL",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "A custom path handler is executed if the request method does not match the route.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "An arbitrary path",
                        "name": "path",
                        "in": "path"
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.AccessToken": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string",
                    "example": "Some kind of JWT"
                }
            }
        },
        "handlers.ErrorMessage": {
            "type": "object",
            "properties": {
                "details": {
                    "description": "Description of the situation",
                    "type": "string",
                    "example": "An error occurred"
                }
            }
        },
        "models.GsmEntryResponse": {
            "type": "object",
            "required": [
                "contractor",
                "dt_receiving",
                "guid",
                "income_kg",
                "license_plate",
                "operator",
                "provider",
                "site",
                "status"
            ],
            "properties": {
                "been_changed": {
                    "description": "The status of the fuel intake record in the database (changed or not)",
                    "type": "boolean",
                    "example": false
                },
                "contractor": {
                    "description": "Name of the fuel carrier",
                    "type": "string",
                    "example": "Name of the fuel carrier"
                },
                "dt_crch": {
                    "description": "Fuel receiving  date",
                    "type": "string",
                    "format": "date",
                    "example": "2025-01-02"
                },
                "dt_receiving": {
                    "description": "Fuel receiving date",
                    "type": "string",
                    "format": "date",
                    "example": "2024-11-15"
                },
                "guid": {
                    "description": "The global unique identifier of the record",
                    "type": "string",
                    "example": "593ff941-405e-4afd-9eec-f8605a14351a"
                },
                "id": {
                    "description": "ID of the database entry",
                    "type": "integer",
                    "minimum": 1
                },
                "income_kg": {
                    "description": "The amount of fuel received at the warehouse in kilograms",
                    "type": "number",
                    "example": 362.2
                },
                "license_plate": {
                    "description": "The state number of the transport that delivered the fuel",
                    "type": "string",
                    "example": " A902RUS"
                },
                "operator": {
                    "description": "Last name of the operator who took the fuel to the warehouse",
                    "type": "string",
                    "example": "Last name of the operator"
                },
                "provider": {
                    "description": "Name of the fuel provider",
                    "type": "string",
                    "example": "Name of the fuel provider"
                },
                "site": {
                    "description": "Name of the mining site",
                    "type": "string",
                    "example": "Name of the mining site"
                },
                "status": {
                    "description": "Fuel loading status",
                    "type": "string",
                    "example": "Uploaded"
                }
            }
        },
        "models.GsmeEntryRequest": {
            "type": "object",
            "required": [
                "contractor",
                "dt_receiving",
                "guid",
                "income_kg",
                "license_plate",
                "operator",
                "provider",
                "site",
                "status"
            ],
            "properties": {
                "been_changed": {
                    "description": "The status of the fuel intake record in the database (changed or not)",
                    "type": "boolean",
                    "example": true
                },
                "contractor": {
                    "description": "Name of the fuel carrier",
                    "type": "string",
                    "example": "Name of the fuel carrier"
                },
                "dt_crch": {
                    "description": "Fuel receiving  date",
                    "type": "string",
                    "format": "date",
                    "example": "2025-01-02"
                },
                "dt_receiving": {
                    "description": "Fuel receiving date",
                    "type": "string",
                    "format": "date",
                    "example": "2024-11-15"
                },
                "guid": {
                    "description": "The global unique identifier of the record",
                    "type": "string",
                    "example": "593ff941-405e-4afd-9eec-f8605a14351a"
                },
                "income_kg": {
                    "description": "The amount of fuel received at the warehouse in kilograms",
                    "type": "number",
                    "example": 362.2
                },
                "license_plate": {
                    "description": "The state number of the transport that delivered the fuel",
                    "type": "string",
                    "example": " A902RUS"
                },
                "operator": {
                    "description": "Last name of the operator who took the fuel to the warehouse",
                    "type": "string",
                    "example": "Last name of the operator"
                },
                "provider": {
                    "description": "Name of the fuel provider",
                    "type": "string",
                    "example": "Name of the fuel provider"
                },
                "site": {
                    "description": "Name of the mining site",
                    "type": "string",
                    "example": "Name of the mining site"
                },
                "status": {
                    "description": "Fuel loading status",
                    "type": "string",
                    "example": "Uploaded"
                }
            }
        },
        "models.IdEntry": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID of the entry in the database",
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "models.UserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "my_password"
                },
                "username": {
                    "type": "string",
                    "example": "Some username"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Enter the JWT token in the format: Bearer Access token",
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
	Title:            "CRM server",
	Description:      "Fuel and Lubricants Accounting Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
