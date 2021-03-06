// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/columns/{columnid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "columns"
                ],
                "summary": "Get column with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Column ID",
                        "name": "columnid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "columns"
                ],
                "summary": "Update column with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Column ID",
                        "name": "columnid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Column info",
                        "name": "column",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Column"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "columns"
                ],
                "summary": "Delete column with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Column ID",
                        "name": "columnid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/columns/{columnid}/tasks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get list of tasks from column with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Column ID",
                        "name": "columnid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create task in column with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Column ID",
                        "name": "columnid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task info",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/comments/{commentid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get comment with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Update comment with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comment text",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete comment with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get Project list",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Create Project",
                "parameters": [
                    {
                        "description": "Project parameters",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/projects/{projectid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get Project with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Update Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Project information",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Delete Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/projects/{projectid}/columns": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "columns"
                ],
                "summary": "Get list of columns in project with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "columns"
                ],
                "summary": "Create column in project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Column info",
                        "name": "column",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Column"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/projects/{projectid}/tasks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get list of tasks for project with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/tasks/{taskid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get task with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update task with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task info",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete task with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/tasks/{taskid}/comments": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get list of comments from task with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Create comment for task with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comment text",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Column": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "pos": {
                    "type": "integer"
                }
            }
        },
        "handlers.Comment": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "handlers.Project": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.Task": {
            "type": "object",
            "properties": {
                "col": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "prio": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Task Management API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
