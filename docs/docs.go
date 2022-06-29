// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/menu/addBaseMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "新增菜单",
                "parameters": [
                    {
                        "description": "菜单信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MenuInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜单信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.MenuResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/menu/deleteBaseMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除菜单",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/menu/getBaseMenuById": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "根据ID获取菜单",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜单信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.MenuResult"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/menu/getBaseMenuTree": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "获取菜单树",
                "responses": {
                    "200": {
                        "description": "菜单树信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.MenuTree"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/menu/getMenuList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "获取菜单列表",
                "responses": {
                    "200": {
                        "description": "菜单列表信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.PageResult"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/menu/updateBaseMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "修改菜单",
                "parameters": [
                    {
                        "description": "菜单信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MenuInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新菜单",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/admin_register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "注册管理员",
                "parameters": [
                    {
                        "description": "用户名, 昵称, 密码, 角色ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户注册账号,返回包括用户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除用户",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "分页获取用户列表",
                "parameters": [
                    {
                        "description": "页码, 每页大小",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分页获取用户列表,返回包括列表,总数,页码,每页数量",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.PageResult"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/resetPassword": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "重置密码",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/setUserInfo": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "修改管理员信息",
                "parameters": [
                    {
                        "description": "ID, 用户名, 昵称, 头像链接",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ChangeUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改管理员信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Menu": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "UpdatedAt": {
                    "type": "string"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                },
                "component": {
                    "description": "前端文件路径",
                    "type": "string"
                },
                "hidden": {
                    "description": "是否在列表隐藏",
                    "type": "boolean"
                },
                "icon": {
                    "description": "图标",
                    "type": "string"
                },
                "name": {
                    "description": "路由名称",
                    "type": "string"
                },
                "parent_id": {
                    "description": "父菜单id",
                    "type": "integer"
                },
                "path": {
                    "description": "路由path",
                    "type": "string"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "title": {
                    "description": "菜单名称",
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "UpdatedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "headerImg": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "request.ById": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "request.ChangeUserInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮件",
                    "type": "string"
                },
                "headerImg": {
                    "description": "头像",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "电话",
                    "type": "string"
                }
            }
        },
        "request.MenuInfo": {
            "type": "object",
            "required": [
                "component",
                "hidden",
                "icon",
                "name",
                "parentId",
                "path",
                "sort",
                "title"
            ],
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "component": {
                    "description": "前端文件路径",
                    "type": "string",
                    "maxLength": 40,
                    "minLength": 1
                },
                "hidden": {
                    "description": "是否隐藏",
                    "type": "boolean"
                },
                "icon": {
                    "description": "icon",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                },
                "name": {
                    "description": "路由名称",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                },
                "parentId": {
                    "description": "父ID",
                    "type": "integer"
                },
                "path": {
                    "description": "路由path",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                },
                "sort": {
                    "description": "排序",
                    "type": "integer",
                    "maximum": 999,
                    "minimum": 1
                },
                "title": {
                    "description": "菜单名称",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                }
            }
        },
        "request.PageInfo": {
            "type": "object",
            "required": [
                "page",
                "pageSize"
            ],
            "properties": {
                "keyword": {
                    "description": "关键字",
                    "type": "string"
                },
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页大小",
                    "type": "integer"
                }
            }
        },
        "request.Register": {
            "type": "object",
            "required": [
                "password",
                "phone",
                "userName"
            ],
            "properties": {
                "email": {
                    "description": "邮件",
                    "type": "string"
                },
                "headerImg": {
                    "description": "头像",
                    "type": "string"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 6
                },
                "phone": {
                    "description": "电话",
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                },
                "userName": {
                    "description": "用户名",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1
                }
            }
        },
        "response.MenuResult": {
            "type": "object",
            "properties": {
                "menu": {
                    "description": "菜单信息",
                    "$ref": "#/definitions/model.Menu"
                }
            }
        },
        "response.MenuTree": {
            "type": "object",
            "properties": {
                "menuList": {
                    "description": "菜单树",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表"
                },
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页大小",
                    "type": "integer"
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "编码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "错误信息"
                }
            }
        },
        "response.UserResult": {
            "type": "object",
            "properties": {
                "user": {
                    "description": "用户信息",
                    "$ref": "#/definitions/model.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GA-API接口文档",
	Description:      "GA管理平台",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
