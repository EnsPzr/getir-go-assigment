// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mehmet Enes PAZAR",
            "url": "https://enespazar.com",
            "email": "m.enespazar@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/records": {
            "get": {
                "description": "This function returns records in database by filters.\nFilter contains startDate, endDate, minCount, maxCount.\nIf startDate has value, createdAt returns records greater than startDate.\nIf endDate has value, createdAt returns records smaller than endDate.\nDate format is YYYY-MM-DD.\nIf minCount has value, sum \"counts\" returns records greater than minCount.\nIf maxCount has value, sum \"counts\" returns records smaller than maxCount.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Get All records by database.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Min Count",
                        "name": "minCount",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Max Count",
                        "name": "maxCount",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "records": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "getir-go-assigment.herokuapp.com",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Getir Go Assigment",
	Description:      "This project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}