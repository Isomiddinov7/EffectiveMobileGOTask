// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/song": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Songs List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Get Songs List",
                "operationId": "get_song_list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetAllSongResponseBody",
                        "schema": {
                            "$ref": "#/definitions/models.GetSongResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Update Song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Update Song",
                "operationId": "update_song",
                "parameters": [
                    {
                        "description": "UpdateSong",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateSong"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Song"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create Song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Create Song",
                "operationId": "create_song",
                "parameters": [
                    {
                        "description": "CreateSongRequestBody",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateSong"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "SongBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Song"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
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
        "/song/{id}": {
            "get": {
                "description": "Get Song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Get Song",
                "operationId": "get_song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetListSongResponseBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Delete Song",
                "operationId": "delete_song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
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
        "handler.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.CreateSong": {
            "type": "object",
            "properties": {
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "group_name": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "lyrics": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song_name": {
                    "type": "string"
                }
            }
        },
        "models.GetSongResponse": {
            "type": "object",
            "properties": {
                "songs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Song"
                    }
                }
            }
        },
        "models.Song": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "group_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "lyrics": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UpdateSong": {
            "type": "object",
            "properties": {
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "group_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "lyrics": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
