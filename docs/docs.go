// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/auth/logout": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "退出登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权管理"
                ],
                "summary": "退出登录",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/api/auth/token": {
            "post": {
                "description": "令牌创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权管理"
                ],
                "summary": "令牌创建",
                "parameters": [
                    {
                        "description": "用户名, 密码",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateAuthTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/api/auth/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权管理"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "get": {
                "description": "服务健康",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础管理"
                ],
                "summary": "服务健康",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/api/user/edit": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "修改基础信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改基础信息",
                "parameters": [
                    {
                        "description": "用户名, 用户呢称",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.EditUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/api/user/{uid}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "获取用户详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateAuthTokenRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "登陆密码",
                    "type": "string"
                },
                "username": {
                    "description": "登陆用户",
                    "type": "string"
                }
            }
        },
        "request.EditUserRequest": {
            "type": "object",
            "required": [
                "nickname",
                "username"
            ],
            "properties": {
                "nickname": {
                    "description": "用户别名",
                    "type": "string"
                },
                "username": {
                    "description": "登录名",
                    "type": "string"
                }
            }
        },
        "response.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务码",
                    "type": "integer"
                },
                "data": {
                    "description": "业务消息体"
                },
                "msg": {
                    "description": "业务消息",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
	Host:        "127.0.0.1:8081",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "gin-skeleton",
	Description: "gin-skeleton 示例项目",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
