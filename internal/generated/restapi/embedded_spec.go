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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Web Application",
    "version": "1.0.0"
  },
  "host": "horse",
  "paths": {
    "/user": {
      "get": {
        "summary": "User list",
        "parameters": [
          {
            "type": "string",
            "name": "fingerprint",
            "in": "header",
            "required": true
          }
        ],
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
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create User",
        "parameters": [
          {
            "type": "string",
            "name": "fingerprint",
            "in": "header",
            "required": true
          },
          {
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/CreateUserRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User was successfully created"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      }
    },
    "/user/{email}": {
      "get": {
        "summary": "Get User by email",
        "parameters": [
          {
            "type": "string",
            "name": "fingerprint",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "email",
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateUserRequest": {
      "description": "Create user request params",
      "type": "object",
      "required": [
        "payload"
      ],
      "properties": {
        "payload": {
          "$ref": "#/definitions/User"
        }
      }
    },
    "FailResponse": {
      "description": "unsuccessful request",
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "description": "error message",
          "type": "string"
        }
      }
    },
    "User": {
      "description": "UserDto",
      "type": "object",
      "required": [
        "email",
        "fullName",
        "isAdmin"
      ],
      "properties": {
        "email": {
          "description": "User email",
          "type": "string",
          "example": "example@gmail.com"
        },
        "fullName": {
          "description": "User full name",
          "type": "string",
          "example": "John Doe"
        },
        "isAdmin": {
          "description": "If User is admin",
          "type": "boolean",
          "example": false
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Web Application",
    "version": "1.0.0"
  },
  "host": "horse",
  "paths": {
    "/user": {
      "get": {
        "summary": "User list",
        "parameters": [
          {
            "type": "string",
            "name": "fingerprint",
            "in": "header",
            "required": true
          }
        ],
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
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Create User",
        "parameters": [
          {
            "type": "string",
            "name": "fingerprint",
            "in": "header",
            "required": true
          },
          {
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/CreateUserRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User was successfully created"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      }
    },
    "/user/{email}": {
      "get": {
        "summary": "Get User by email",
        "parameters": [
          {
            "type": "string",
            "name": "fingerprint",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "email",
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateUserRequest": {
      "description": "Create user request params",
      "type": "object",
      "required": [
        "payload"
      ],
      "properties": {
        "payload": {
          "$ref": "#/definitions/User"
        }
      }
    },
    "FailResponse": {
      "description": "unsuccessful request",
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "description": "error message",
          "type": "string"
        }
      }
    },
    "User": {
      "description": "UserDto",
      "type": "object",
      "required": [
        "email",
        "fullName",
        "isAdmin"
      ],
      "properties": {
        "email": {
          "description": "User email",
          "type": "string",
          "example": "example@gmail.com"
        },
        "fullName": {
          "description": "User full name",
          "type": "string",
          "example": "John Doe"
        },
        "isAdmin": {
          "description": "If User is admin",
          "type": "boolean",
          "example": false
        }
      }
    }
  }
}`))
}
