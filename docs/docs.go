// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Eraspace",
            "url": "eraspace.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/match/recommendations/user/{id}/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Endpoint for get list recommendations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Accept",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Content-Type",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Example: 5d47eb91-bee9-46b8-9104-aea0f40ef1c3",
                        "name": "Device-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Example: eraspace",
                        "name": "Source",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "name": "age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "0",
                        "description": "Optional, will fill with default value 0",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "CursorDirectionNext",
                            "CursorDirectionPrev"
                        ],
                        "description": "Optional, will fill with default value NEXT",
                        "name": "cursorDir",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "Optional, will fill with default value 10",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "id"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "CustomerSortByID"
                        ],
                        "description": "\"id\" is the same as \"created at\"",
                        "name": "sortBy",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "CustomerSortDirAscending",
                            "CustomerSortDirDescending"
                        ],
                        "description": "Default value is asc",
                        "name": "sortDir",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entity.SwaggerResponseOKDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/http.cursorPaginationResponse-entity_Users"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/entity.Users"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/entity.CursorInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "*Notes: Code data will be return null",
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerResponseBadRequestDTO"
                        }
                    },
                    "401": {
                        "description": "*Notes: Code data will be return null",
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerResponseUnauthorizedDTO"
                        }
                    },
                    "500": {
                        "description": "*Notes: Code data will be return null",
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerResponseInternalServerErrorDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.CursorDirection": {
            "type": "string",
            "enum": [
                "next",
                "prev"
            ],
            "x-enum-varnames": [
                "CursorDirectionNext",
                "CursorDirectionPrev"
            ]
        },
        "entity.CursorInfo": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 20
                },
                "cursor": {
                    "type": "string",
                    "example": "1696466522533538969"
                },
                "cursorType": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.CursorDirection"
                        }
                    ],
                    "example": "next"
                },
                "hasNext": {
                    "type": "boolean",
                    "example": true
                },
                "hasPrev": {
                    "type": "boolean",
                    "example": true
                },
                "nextCursor": {
                    "type": "string",
                    "example": "1695785802835854036"
                },
                "prevCursor": {
                    "type": "string",
                    "example": "1696415865308136181"
                },
                "size": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "entity.CustomerSortDir": {
            "type": "string",
            "enum": [
                "asc",
                "desc"
            ],
            "x-enum-varnames": [
                "CustomerSortDirAscending",
                "CustomerSortDirDescending"
            ]
        },
        "entity.CustomerURLSortBy": {
            "type": "string",
            "enum": [
                "id"
            ],
            "x-enum-varnames": [
                "CustomerSortByID"
            ]
        },
        "entity.SwaggerResponseBadRequestDTO": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string",
                    "example": "Customer Miscellaneous API"
                },
                "build": {
                    "type": "string",
                    "example": "1"
                },
                "data": {
                    "description": "Will return null"
                },
                "id": {
                    "type": "string",
                    "example": "16ad78a0-5f8a-4af0-9946-d21656e718b5"
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                }
            }
        },
        "entity.SwaggerResponseInternalServerErrorDTO": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string",
                    "example": "Customer Miscellaneous API"
                },
                "build": {
                    "type": "string",
                    "example": "1"
                },
                "data": {
                    "description": "Will return null"
                },
                "id": {
                    "type": "string",
                    "example": "16ad78a0-5f8a-4af0-9946-d21656e718b5"
                },
                "message": {
                    "type": "string",
                    "example": "Internal Server Error"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                }
            }
        },
        "entity.SwaggerResponseOKDTO": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string",
                    "example": "Customer Miscellaneous API"
                },
                "build": {
                    "type": "string",
                    "example": "1"
                },
                "data": {},
                "id": {
                    "type": "string",
                    "example": "16ad78a0-5f8a-4af0-9946-d21656e718b5"
                },
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                }
            }
        },
        "entity.SwaggerResponseUnauthorizedDTO": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string",
                    "example": "Customer Miscellaneous API"
                },
                "build": {
                    "type": "string",
                    "example": "1"
                },
                "data": {
                    "description": "Will return null"
                },
                "id": {
                    "type": "string",
                    "example": "16ad78a0-5f8a-4af0-9946-d21656e718b5"
                },
                "message": {
                    "type": "string",
                    "example": "Unauthorized"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                }
            }
        },
        "entity.Users": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "interest": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "preferences": {
                    "type": "string"
                }
            }
        },
        "http.cursorPaginationResponse-entity_Users": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Users"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/entity.CursorInfo"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
