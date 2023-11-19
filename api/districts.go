package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetDistricts godoc
// @Description Get's all districts from civ 6
// @Tags Districts
// @Accept */*
// @Produce json
// @Param limit query int false "limits amount of results returned"
// @Param civilization query string false "filters leaders by civilization"
// @Success 200 {object} []database.District
// @Failure 400 {object} string "Invalid limit value"
// @Router /api/v1/districts [get]
func (r *Router) GetDistricts(c echo.Context) error {
	districts := r.db.Districts

	limit := c.QueryParam("limit")

	if limit == "" {
		return c.JSON(http.StatusOK, districts)
	} else {
		limit, err := CheckLimit(limit)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, districts[:limit])
	}
}

// GetDistrict godoc
// @Description Gets a district by name in civ 6
// @Tags Districts
// @Accept */*
// @Produce json
// @Param name path string true "district to find"
// @Success 200 {object} database.District
// @Router /api/v1/districts/{name} [get]
func (r *Router) GetDistrict(c echo.Context) error {
	name := c.Param("name")
	name = strings.ReplaceAll(name, "%20", " ")

	for _, district := range r.db.Districts {
		if strings.ToLower(district.Name) == strings.ToLower(name) {
			return c.JSON(http.StatusOK, district)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Could not find civilzation")
}
