package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetCivlizations(c echo.Context) error {
	civs := r.db.Civilizations

	limit := c.QueryParam("limit")

	if limit == "" {
		return c.JSON(http.StatusOK, civs)
	} else {
		limit, err := strconv.Atoi(limit)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit parameter")
		}

		return c.JSON(http.StatusOK, civs[:limit])
	}
}

func (r *Router) GetCivilzation(c echo.Context) error {
	name := c.Param("name")
	name = strings.ReplaceAll(name, "%20", " ")

	for _, civ := range r.db.Civilizations {
		if strings.ToLower(civ.Name) == strings.ToLower(name) {
			return c.JSON(http.StatusOK, civ)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Could not find civilzation")
}
