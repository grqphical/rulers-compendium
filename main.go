package main

import (
	"example/civ6-api/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}

func get_leaders(c echo.Context) error {
	leaders := database.ReadLeaders()

}

func main() {
	e := echo.New()

	e.Static("/static", "./static")

	e.GET("/", index)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
