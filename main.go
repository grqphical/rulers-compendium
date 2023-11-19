package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/grqphical07/rulers-compendium/api"
	"github.com/grqphical07/rulers-compendium/database"
	_ "github.com/grqphical07/rulers-compendium/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Rulers Compendium API
// @version 1.0
// @description A free-to-use API to access information about Sid Meier's Civilization VI
// @termsOfService https://github.com/grqphical07/rulers-compendium/blob/main/TERMS.md

// @contact.name grqphical
// @contact.url https://github.com/grqphical07

// @license.name MIT License
// @license.url https://github.com/grqphical07/rulers-compendium/blob/main/LICENSE

// @host localhost:8000
// @BasePath /api/v1
// @schemes http
func main() {
	e := echo.New()

	db := database.ReadDatabase()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%v - %v\n", v.URI, v.Status)
			return nil
		},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.CORS())

	api_routes := e.Group("/api")

	version_1 := api_routes.Group("/v1")

	router := api.NewRouter(&db, e)

	// Leaders API
	version_1.GET("/", router.Index)
	version_1.GET("/docs/*", echoSwagger.WrapHandler)

	leaders_api := version_1.Group("/leaders")
	leaders_api.GET("", router.GetLeaders)
	leaders_api.GET("/:name", router.GetLeader)

	// Civilizations API
	civ_api := version_1.Group("/civilizations")
	civ_api.GET("", router.GetCivlizations)
	civ_api.GET("/:name", router.GetCivilzation)

	// Districts API
	districts_api := version_1.Group("/districts")
	districts_api.GET("", router.GetDistricts)
	districts_api.GET("/:name", router.GetDistrict)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
