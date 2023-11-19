package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetCivilizations godoc
// @Description Get's all civilizations available to play in civ 6
// @Tags leaders
// @Accept */*
// @Produce json
// @Param limit query int false "limits amount of results returned"
// @Success 200 {object} []database.Civilization
// @Failure 400 {object} string "Invalid limit value"
// @Router /api/v1/civilizations [get]
func (r *Router) GetCivilizations(c echo.Context) error {
	civs := r.db.Civilizations

	limit := c.QueryParam("limit")

	if limit == "" {
		return c.JSON(http.StatusOK, civs)
	} else {
		limit, err := CheckLimit(limit)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, civs[:limit])
	}
}

// GetCivilization godoc
// @Description Gets a civilization in civ 6 based on a given name
// @Tags leaders
// @Accept */*
// @Produce json
// @Param name path string true "civilization to get"
// @Success 200 {object} database.Civilization
// @Router /api/v1/civilizations/{name} [get]
func (r *Router) GetCivilization(c echo.Context) error {
	name := c.Param("name")
	name = strings.ReplaceAll(name, "%20", " ")

	for _, civ := range r.db.Civilizations {
		if strings.ToLower(civ.Name) == strings.ToLower(name) {
			return c.JSON(http.StatusOK, civ)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Could not find civilzation")
}
