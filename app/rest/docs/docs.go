// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Coding4u",
            "email": "oi@coding4u.com.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/guests": {
            "post": {
                "description": "Router for create a new guest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Guest"
                ],
                "summary": "create a new guest",
                "operationId": "createGuest",
                "parameters": [
                    {
                        "description": "JSON body for create a new guest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.CreateGuestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/guests/{guest_id}": {
            "get": {
                "description": "Router for find a guest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Guest"
                ],
                "summary": "find a gust",
                "operationId": "findGuest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Guest ID",
                        "name": "guest_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Guest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/scores": {
            "post": {
                "description": "Router for create a new score",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Score"
                ],
                "summary": "create a new score",
                "operationId": "createScore",
                "parameters": [
                    {
                        "description": "JSON body for create a new score",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.CreateScoreRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/scores/{score_id}": {
            "get": {
                "description": "Router for find a score",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Score"
                ],
                "summary": "find a score",
                "operationId": "findScore",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Score ID",
                        "name": "score_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Score"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/scores/{score_id}/tags": {
            "post": {
                "description": "Router for add a tag to a score",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Score"
                ],
                "summary": "add a tag to a score",
                "operationId": "addTag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Score ID",
                        "name": "score_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "JSON body for add a tag to a score",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.AddTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/scores/{score_id}/use": {
            "post": {
                "description": "Router for use a score",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Score"
                ],
                "summary": "use a score",
                "operationId": "useScore",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Score ID",
                        "name": "score_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "description": "Router for search tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "search tags",
                "operationId": "searchTags",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/rest.Tag"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Router for create a new tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "create a new tag",
                "operationId": "createTag",
                "parameters": [
                    {
                        "description": "JSON body for create a new tag",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.CreateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/tags/{tag_id}": {
            "get": {
                "description": "Router for find a tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "find a tag",
                "operationId": "findTag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag ID",
                        "name": "tag_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Tag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/rest.HTTPResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "rest.AddTagRequest": {
            "type": "object",
            "properties": {
                "tag_id": {
                    "type": "string"
                }
            }
        },
        "rest.CreateGuestRequest": {
            "type": "object",
            "properties": {
                "doc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "rest.CreateScoreRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "guest_id": {
                    "type": "string"
                }
            }
        },
        "rest.CreateTagRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "rest.Guest": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "doc": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "rest.HTTPResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "any message"
                }
            }
        },
        "rest.IDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "rest.Score": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "doc": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/rest.Tag"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "used_in": {
                    "type": "string"
                },
                "was_used": {
                    "type": "boolean"
                }
            }
        },
        "rest.Tag": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Loyalty Card Swagger API",
	Description: "Swagger API for Loyalty Card Service.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
