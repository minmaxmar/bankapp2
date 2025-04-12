// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "bankapp2",
    "title": "bankapp2",
    "version": "2.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/v2",
  "paths": {
    "/banks": {
      "get": {
        "summary": "Get list of banks",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Bank"
              }
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create a new bank",
        "parameters": [
          {
            "description": "Bank to be created",
            "name": "bank",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewBank"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Bank created",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "patch": {
        "summary": "Patch bank",
        "parameters": [
          {
            "description": "Bank patched",
            "name": "bank",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Bank patched",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/banks/{id}": {
      "get": {
        "summary": "Get a bank by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "404": {
            "description": "Bank not found"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a bank by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Bank deleted"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/cards": {
      "get": {
        "summary": "Get list of cards",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Card"
              }
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create a new card",
        "parameters": [
          {
            "description": "Card to be created",
            "name": "card",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewCard"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Card created",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "patch": {
        "summary": "Patch card",
        "parameters": [
          {
            "description": "Card to be patched",
            "name": "card",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Card"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Card patched",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/cards/{id}": {
      "get": {
        "summary": "Get a card by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "404": {
            "description": "Card not found"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a card by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Card deleted"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "summary": "Get list of users",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create a new user",
        "parameters": [
          {
            "description": "User to be created",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewUser"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "patch": {
        "summary": "Patch user",
        "parameters": [
          {
            "description": "User patched",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User patched",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "summary": "Get a user by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a user by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "User deleted"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Bank": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        }
      }
    },
    "Card": {
      "type": "object",
      "properties": {
        "BankID": {
          "type": "integer"
        },
        "CreateDate": {
          "type": "string",
          "format": "date-time"
        },
        "ExpiryDate": {
          "type": "string",
          "format": "date"
        },
        "Number": {
          "type": "integer"
        },
        "Total": {
          "type": "integer"
        },
        "UserID": {
          "type": "integer"
        },
        "id": {
          "type": "integer"
        }
      }
    },
    "ErrorResponse": {
      "description": "Общая ошибка",
      "title": "ErrorResponse",
      "allOf": [
        {
          "properties": {
            "error": {
              "type": "object",
              "properties": {
                "message": {
                  "description": "Message",
                  "type": "string"
                }
              }
            }
          }
        }
      ]
    },
    "NewBank": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        }
      }
    },
    "NewCard": {
      "type": "object",
      "properties": {
        "BankID": {
          "type": "integer",
          "x-go-custom-tag": "validate:\"required\""
        },
        "CreateDate": {
          "type": "string",
          "format": "date-time"
        },
        "ExpiryDate": {
          "type": "string",
          "format": "date",
          "x-go-custom-tag": "validate:\"required\""
        },
        "Number": {
          "type": "integer",
          "x-go-custom-tag": "validate:\"required\""
        },
        "Total": {
          "type": "integer"
        },
        "UserID": {
          "type": "integer",
          "x-go-custom-tag": "validate:\"required\""
        }
      }
    },
    "NewUser": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        },
        "firstName": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        },
        "lastName": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "lastName": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "bankapp2",
    "title": "bankapp2",
    "version": "2.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/v2",
  "paths": {
    "/banks": {
      "get": {
        "summary": "Get list of banks",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Bank"
              }
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create a new bank",
        "parameters": [
          {
            "description": "Bank to be created",
            "name": "bank",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewBank"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Bank created",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "patch": {
        "summary": "Patch bank",
        "parameters": [
          {
            "description": "Bank patched",
            "name": "bank",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Bank patched",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/banks/{id}": {
      "get": {
        "summary": "Get a bank by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "404": {
            "description": "Bank not found"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a bank by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Bank deleted"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/cards": {
      "get": {
        "summary": "Get list of cards",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Card"
              }
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create a new card",
        "parameters": [
          {
            "description": "Card to be created",
            "name": "card",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewCard"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Card created",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "patch": {
        "summary": "Patch card",
        "parameters": [
          {
            "description": "Card to be patched",
            "name": "card",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Card"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Card patched",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/cards/{id}": {
      "get": {
        "summary": "Get a card by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "404": {
            "description": "Card not found"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a card by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Card deleted"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "summary": "Get list of users",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create a new user",
        "parameters": [
          {
            "description": "User to be created",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewUser"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "patch": {
        "summary": "Patch user",
        "parameters": [
          {
            "description": "User patched",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User patched",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "summary": "Get a user by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a user by ID",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "User deleted"
          },
          "default": {
            "description": "Общая ошибка",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Bank": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        }
      }
    },
    "Card": {
      "type": "object",
      "properties": {
        "BankID": {
          "type": "integer"
        },
        "CreateDate": {
          "type": "string",
          "format": "date-time"
        },
        "ExpiryDate": {
          "type": "string",
          "format": "date"
        },
        "Number": {
          "type": "integer"
        },
        "Total": {
          "type": "integer"
        },
        "UserID": {
          "type": "integer"
        },
        "id": {
          "type": "integer"
        }
      }
    },
    "ErrorResponse": {
      "description": "Общая ошибка",
      "title": "ErrorResponse",
      "allOf": [
        {
          "properties": {
            "error": {
              "type": "object",
              "properties": {
                "message": {
                  "description": "Message",
                  "type": "string"
                }
              }
            }
          }
        }
      ]
    },
    "ErrorResponseAO0Error": {
      "type": "object",
      "properties": {
        "message": {
          "description": "Message",
          "type": "string"
        }
      }
    },
    "NewBank": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        }
      }
    },
    "NewCard": {
      "type": "object",
      "properties": {
        "BankID": {
          "type": "integer",
          "x-go-custom-tag": "validate:\"required\""
        },
        "CreateDate": {
          "type": "string",
          "format": "date-time"
        },
        "ExpiryDate": {
          "type": "string",
          "format": "date",
          "x-go-custom-tag": "validate:\"required\""
        },
        "Number": {
          "type": "integer",
          "x-go-custom-tag": "validate:\"required\""
        },
        "Total": {
          "type": "integer"
        },
        "UserID": {
          "type": "integer",
          "x-go-custom-tag": "validate:\"required\""
        }
      }
    },
    "NewUser": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        },
        "firstName": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        },
        "lastName": {
          "type": "string",
          "x-go-custom-tag": "validate:\"required\""
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "lastName": {
          "type": "string"
        }
      }
    }
  }
}`))
}
