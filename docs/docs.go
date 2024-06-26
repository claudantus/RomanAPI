// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
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
        "/": {
            "get": {
                "description": "Displays a short description of how to use the API",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Shows welcome page",
                "responses": {
                    "200": {
                        "description": "Welcome message",
                        "schema": {
                            "$ref": "#/definitions/api.welcomeMessage"
                        }
                    }
                }
            }
        },
        "/romans": {
            "get": {
                "description": "Uses a min and a max parameter to define the range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Returns Roman numerals in a specified range",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Min",
                        "name": "min",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Max",
                        "name": "max",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.decimalRoman"
                            }
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.ErrorMsg"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrorMsg": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "Min"
                },
                "message": {
                    "type": "string",
                    "example": "This field is required and needs to be an integer in the range of 1-3999"
                }
            }
        },
        "api.decimalRoman": {
            "type": "object",
            "properties": {
                "decimal": {
                    "type": "integer",
                    "example": 10
                },
                "roman": {
                    "type": "string",
                    "example": "X"
                }
            }
        },
        "api.welcomeMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Welcome to the Roman Numeral API.\nGet a range of roman numerals via /api/v1/romans with the query parameters \n\"min\" for the lower and\n\"max\" for the upper bound"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Roman Numeral API",
	Description:      "Returns roman numerals in a range specified by query parameters",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
