package main

import (
	"fmt"
	"net/http"

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

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/api/v1/docs/index.html")
	})

	api_routes := e.Group("/api")

	version_1 := api_routes.Group("/v1")
	version_1.GET("/docs/*", echoSwagger.WrapHandler)

	router := api.NewRouter(&db, e)

	// Leaders API
	leaders_api := version_1.Group("/leaders")
	leaders_api.GET("", router.GetLeaders)
	leaders_api.GET("/:name", router.GetLeader)

	// Civilizations API
	civ_api := version_1.Group("/civilizations")
	civ_api.GET("", router.GetCivilizations)
	civ_api.GET("/:name", router.GetCivilization)

	// Districts API
	districts_api := version_1.Group("/districts")
	districts_api.GET("", router.GetDistricts)
	districts_api.GET("/:name", router.GetDistrict)

	// Improvements API
	improvements_api := version_1.Group("/improvements")
	improvements_api.GET("", router.GetImprovements)
	improvements_api.GET("/:name", router.GetImprovement)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
