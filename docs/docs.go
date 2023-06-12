// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Marcos Coutinho",
            "url": "https://linkedin.com/in/marcoscoutinhodev",
            "email": "marcoscoutinhodev@outlook.com"
        },
        "license": {
            "name": "The MIT License (MIT)",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/url": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "url"
                ],
                "summary": "Create URL",
                "parameters": [
                    {
                        "description": "url request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.ShortURLInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONError"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "Authenticate User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Authenticate User",
                "parameters": [
                    {
                        "description": "user request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.UserInputSignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONError"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "user request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.UserInputSignUp"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ToJSONError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "swagger.ShortURLInput": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "swagger.ToJSONError": {
            "type": "object",
            "properties": {
                "error": {},
                "success": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "swagger.ToJSONSuccess": {
            "type": "object",
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean"
                }
            }
        },
        "swagger.UserInputSignIn": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "swagger.UserInputSignUp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-access-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:4001",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "URL SHORTENER API",
	Description:      "api for url shortener application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}