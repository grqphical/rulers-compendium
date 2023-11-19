package api

import (
	"net/http"
	"strconv"

	"github.com/grqphical07/rulers-compendium/database"

	"github.com/labstack/echo/v4"
)

// Manages state between routes in the API
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

// Checks if a limit query parameter is valid meaning it's an integer above zero
// If it isn't it generates an HTTP error to be sent to the client
func CheckLimit(limit string) (int, error) {
	limit_int, err := strconv.Atoi(limit)
	if err != nil || limit_int <= 0 {
		return -1, echo.NewHTTPError(http.StatusBadRequest, "Invalid limit parameter. Limit must be integer above zero")
	}
	return limit_int, nil
}
