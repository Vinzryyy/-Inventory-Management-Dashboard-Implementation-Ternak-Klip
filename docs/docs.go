package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "title": "Inventory Management Dashboard API",
        "description": "Backend API untuk inventory dashboard dengan Echo, PostgreSQL/Supabase, dan Swagger.",
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "schemes": ["http"],
    "paths": {
        "/products": {
            "get": {
                "produces": ["application/json"],
                "summary": "List products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Product": {
            "type": "object",
            "properties": {
                "id": {"type": "integer"},
                "name": {"type": "string"},
                "stock_count": {"type": "integer"},
                "last_updated": {"type": "string"}
            }
        },
        "models.ProductListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct{}

func (s *swaggerInfo) ReadDoc() string {
	return docTemplate
}

func init() {
	swag.Register(swag.Name, &swaggerInfo{})
}
