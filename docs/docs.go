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
        "contact": {
            "name": "lily-lee",
            "email": "lilylee99.01@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/drafts": {
            "get": {
                "description": "list drafts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drafts"
                ],
                "summary": "list drafts",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "volumeID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/draft.ListResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            },
            "post": {
                "description": "create a draft article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drafts"
                ],
                "summary": "create a draft article",
                "parameters": [
                    {
                        "description": "create param",
                        "name": "draft",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/draft.CreateParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Draft"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/drafts/{id}": {
            "get": {
                "description": "get a specific article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drafts"
                ],
                "summary": "get a specific article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "draft id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Draft"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            },
            "put": {
                "description": "edit an article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drafts"
                ],
                "summary": "edit an article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "draft id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create param",
                        "name": "draft",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/draft.EditParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Draft"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/drafts/{id}/posts": {
            "post": {
                "description": "edit an article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drafts"
                ],
                "summary": "edit an article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "draft id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Draft"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "user login param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.RegisterResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/posts": {
            "get": {
                "description": "list posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "list posts",
                "parameters": [
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "userID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "volumeID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.ListResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/posts/{id}": {
            "get": {
                "description": "return one post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "get a specific post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "post id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/posts/{id}/comments": {
            "post": {
                "description": "comment a post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "comment a post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "post id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create comment param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.CreateCommentParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/posts/{id}/like": {
            "post": {
                "description": "When you request this api if you had like a post, then you'll unlike the post.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "like or dislike post.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "post id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "login token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.LikeResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        },
        "/api/registration": {
            "post": {
                "description": "register a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "user register",
                "parameters": [
                    {
                        "description": "user register param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.RegisterResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/request.BizErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "draft.CreateParam": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_url": {
                    "type": "string"
                },
                "digest": {
                    "type": "string"
                },
                "tag": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "volume_id": {
                    "type": "integer"
                }
            }
        },
        "draft.EditParam": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_url": {
                    "type": "string"
                },
                "digest": {
                    "type": "string"
                },
                "tag": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "volume_id": {
                    "type": "integer"
                }
            }
        },
        "draft.ListResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Draft"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                }
            }
        },
        "jwttoken.Token": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "integer"
                },
                "issued_at": {
                    "type": "integer"
                },
                "signature": {
                    "type": "string"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "properties": {
                "anonymous": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "pid": {
                    "type": "integer"
                },
                "post_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Draft": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "digest": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "post_id": {
                    "type": "integer"
                },
                "posted": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "volume_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "post.CreateCommentParam": {
            "type": "object",
            "required": [
                "content"
            ],
            "properties": {
                "anonymous": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer"
                }
            }
        },
        "post.Item": {
            "type": "object",
            "properties": {
                "comment_count": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "cover_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "digest": {
                    "type": "string"
                },
                "draft_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "like_count": {
                    "type": "integer"
                },
                "tag": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "volume_id": {
                    "type": "integer"
                }
            }
        },
        "post.LikeResp": {
            "type": "object",
            "properties": {
                "like": {
                    "type": "boolean"
                },
                "post_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "post.ListResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Item"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                }
            }
        },
        "post.Post": {
            "type": "object",
            "properties": {
                "comment_count": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "cover_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "digest": {
                    "type": "string"
                },
                "draft_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "like_count": {
                    "type": "integer"
                },
                "liked": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "volume_id": {
                    "type": "integer"
                }
            }
        },
        "request.BizErr": {
            "type": "object",
            "properties": {
                "err_msg": {
                    "type": "string"
                }
            }
        },
        "user.LoginParam": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.RegisterParam": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.RegisterResp": {
            "type": "object",
            "properties": {
                "token": {
                    "$ref": "#/definitions/jwttoken.Token"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "blog api server",
	Description:      "This is a blog api server, serves api for blog web frontend.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
