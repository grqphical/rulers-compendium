package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}

func main() {
	e := echo.New()

	e.Static("/static", "./static")

	e.GET("/", index)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
