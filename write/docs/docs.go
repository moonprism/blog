// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-03 16:42:35.136108823 +0800 CST m=+11.440766434

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/article": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "current page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.articleListResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "create article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "articleInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ArticleInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/article/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "show article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleDetail"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "update article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "articleUpdateBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.ArticleUpdateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "delete article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "admin login (jwt)",
                "parameters": [
                    {
                        "description": " ",
                        "name": "loginData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.AuthBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.authResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/code": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "get code list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "current page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.codeListResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "create code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "codeMeta",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.CodeUpdateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/code/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "show code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "code id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CodeDetail"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "update code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "code id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "codeUpdateBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.CodeUpdateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "delete code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "code id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/file/image": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "get file list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "current time yyyy-MM-dd HH:mm:ss",
                        "name": "max_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "fetch list length",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.fileListResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "upload image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": " ",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FileMeta"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/file/image/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "delete file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/tag": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "get tag list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TagInfo"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "create tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "tagMeta",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.TagMeta"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/tag/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "update tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "tag id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "TagMeta",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.TagMeta"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "delete tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "tag id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthenticated",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.articleListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ArticleInfo"
                    }
                },
                "pagination": {
                    "type": "object",
                    "$ref": "#/definitions/utils.Pagination"
                }
            }
        },
        "api.authResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "api.codeListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CodeMeta"
                    }
                },
                "pagination": {
                    "type": "object",
                    "$ref": "#/definitions/utils.Pagination"
                }
            }
        },
        "api.fileListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.File"
                    }
                }
            }
        },
        "errors.HTTPError": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.ArticleDetail": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "html": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TagMeta"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "model.ArticleInfo": {
            "type": "object",
            "properties": {
                "created_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TagMeta"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "model.CodeDetail": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lang": {
                    "type": "string"
                },
                "tags": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "model.CodeMeta": {
            "type": "object",
            "properties": {
                "created_time": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lang": {
                    "type": "string"
                },
                "tags": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "model.File": {
            "type": "object",
            "properties": {
                "created_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "model.FileMeta": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                }
            }
        },
        "model.TagInfo": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.TagMeta": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.ArticleUpdateBody": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "html": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TagMeta"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "service.AuthBody": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.CodeUpdateBody": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lang": {
                    "type": "string"
                },
                "tags": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "utils.Pagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
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
	Title:       "Kicoe Blog API",
	Description: "This is blog api test server. jwt token prefix: Bearer",
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
