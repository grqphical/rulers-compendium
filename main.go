package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"example/civ6-api/api"
	"example/civ6-api/database"
)

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

	api_routes := e.Group("/api")

	version_1 := api_routes.Group("/v1")

	router := api.NewRouter(&db, e)

	// Leaders API
	version_1.GET("/", router.Index)
	version_1.GET("/leaders", router.GetLeaders)
	version_1.GET("/leaders/:name", router.GetLeader)

	// Civilizations API
	version_1.GET("/civilizations", router.GetCivlizations)
	version_1.GET("/civilizations/:name", router.GetCivilzation)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
