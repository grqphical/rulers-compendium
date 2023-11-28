package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetImprovements godoc
// @Description Get's all improvements buildable in civ 6
// @Tags Improvements
// @Accept */*
// @Produce json
// @Param limit query int false "limits amount of results returned"
// @Success 200 {object} []database.Improvement
// @Failure 400 {object} string "Invalid limit value"
// @Router /improvements [get]
func (r *Router) GetImprovements(c echo.Context) error {
	limit := c.QueryParam("limit")
	improvements := r.db.Improvements

	if limit == "" {
		return c.JSON(http.StatusOK, improvements)
	} else {
		limit, err := CheckLimit(limit)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, improvements[:limit])
	}
}

// GetImprovement godoc
// @Description Gets an improvement in civ 6 based on a given name
// @Tags Improvements
// @Accept */*
// @Produce json
// @Param name path string true "Improvement to get"
// @Success 200 {object} database.Improvement
// @Router /improvements/{name} [get]
func (r *Router) GetImprovement(c echo.Context) error {
	name := c.Param("name")
	name = strings.ReplaceAll(name, "%20", " ")

	for _, improvement := range r.db.Improvements {
		if strings.ToLower(improvement.Name) == strings.ToLower(name) {
			return c.JSON(http.StatusOK, improvement)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Could not find improvement")
}
