package main

import (
	"log"

	"inventory-dashboard/database"
	_ "inventory-dashboard/docs"
	"inventory-dashboard/repo"
	"inventory-dashboard/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Inventory Management Dashboard API
// @version 1.0
// @description Backend API untuk inventory dashboard menggunakan Echo, PostgreSQL/Supabase, dan Swaggo.
// @BasePath /api/v1
func main() {
	cfg, err := database.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := NewHandler(service.NewProductService(repo.NewProductRepository(db)))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.AllowedOrigin},
		AllowMethods: []string{echo.GET, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Static("/assets", "frontend")
	e.File("/", "frontend/index.html")
	e.GET("/healthz", handler.Health)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api/v1")
	api.GET("/products", handler.ListProducts)

	log.Fatal(e.Start(":" + cfg.Port))
}
