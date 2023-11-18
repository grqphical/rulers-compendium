package api

import (
	"example/civ6-api/database"

	"github.com/labstack/echo/v4"
)

type Router struct {
	engine *echo.Echo
	db     *database.Database
}

func NewRouter(db *database.Database, e *echo.Echo) Router {
	return Router{
		engine: e,
		db:     db,
	}
}
