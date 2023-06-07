// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/user": {
            "put": {
                "description": "유저 정보를 등록합니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "유저 정보 관리"
                ],
                "summary": "유저 정보 등록",
                "parameters": [
                    {
                        "description": "유저 정보 JSON Format",
                        "name": "Param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserResult"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "유저 정보를 등록합니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "유저 정보 관리"
                ],
                "summary": "유저 정보 등록",
                "parameters": [
                    {
                        "description": "유저 정보 JSON Format",
                        "name": "Param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserResult"
                            }
                        }
                    }
                }
            }
        },
        "/user/": {
            "get": {
                "description": "정산 정보 목록을 가져옵니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "유저 정보 관리"
                ],
                "summary": "정산 정보 목록 열람",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserListResult"
                            }
                        }
                    }
                }
            }
        },
        "/user/{userId}": {
            "get": {
                "description": "user_id과 관련된 유저 정보값을 가져옵니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "유저 정보 관리"
                ],
                "summary": "유저 정보 열람",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "가져올 유저 정보의 유저 시리얼(아이디)",
                        "name": "userId",
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
                                "$ref": "#/definitions/model.UserResult"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "user_id과 관련된 유저 정보값을 가져옵니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "유저 정보 관리"
                ],
                "summary": "유저 정보 열람",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "삭제할 유저 정보의 유저 시리얼(아이디)",
                        "name": "accountSrl",
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
                                "$ref": "#/definitions/model.UserResult"
                            }
                        }
                    }
                }
            }
        },
        "/welcome/{name}": {
            "get": {
                "description": "스웨거 테스트용 핸들러 Desc",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Welcome!"
                ],
                "summary": "스웨거 테스트용 핸들러",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.welcomeModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "f_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "l_name": {
                    "type": "string"
                }
            }
        },
        "model.UserListResult": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "string",
                    "example": "INSERT/SELECT/UPDATE/DELETE"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                },
                "user_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "model.UserResult": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "string",
                    "example": "INSERT/SELECT/UPDATE/DELETE"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                },
                "user_data": {
                    "$ref": "#/definitions/model.User"
                }
            }
        },
        "router.welcomeModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "account name"
                }
            }
        }
    },
    "securityDefinitions": {
        "user-token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GO BASE PROJ",
	Description:      "GO BASE PROJ API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
