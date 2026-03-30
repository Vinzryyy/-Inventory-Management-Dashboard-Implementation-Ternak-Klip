package main

import (
	"fmt"
	"net/http"

	"inventory-dashboard/models"
	"inventory-dashboard/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	productService service.ProductService
}

func NewHandler(productService service.ProductService) *Handler {
	return &Handler{
		productService: productService,
	}
}

// Health godoc
// @Summary Health check
// @Tags Health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /healthz [get]
func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthResponse{Status: "ok"})
}

// ListProducts godoc
// @Summary List inventory products
// @Tags Products
// @Produce json
// @Success 200 {object} models.ProductListResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products [get]
func (h *Handler) ListProducts(c echo.Context) error {
	products, err := h.productService.ListProducts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: fmt.Sprintf("failed to fetch products: %v", err)})
	}

	return c.JSON(http.StatusOK, models.ProductListResponse{Data: products})
}
