package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetWonders godoc
// @Description Get's all buildable wonders from civ 6
// @Tags Wonders
// @Accept */*
// @Produce json
// @Param limit query int false "limits amount of results returned"
// @Success 200 {object} []database.Wonder
// @Failure 400 {object} string "Invalid limit value"
// @Router /wonders [get]
func (r *Router) GetWonders(c echo.Context) error {
	limit := c.QueryParam("limit")

	if limit == "" {
		wonders := r.db.Wonders
		return c.JSON(http.StatusOK, wonders)
	} else {
		limit, err := CheckLimit(limit)
		if err != nil {
			return err
		}

		wonders := r.db.Wonders[:limit]
		return c.JSON(http.StatusOK, wonders)
	}

}

// GetWonder godoc
// @Description Gets a buildable wonder from civ 6
// @Tags Wonders
// @Accept */*
// @Produce json
// @Param name path string true "wonder to get"
// @Success 200 {object} database.Wonder
// @Router /wonders/{name} [get]
func (r *Router) GetWonder(c echo.Context) error {
	name := c.Param("name")
	name = strings.ReplaceAll(name, "%20", " ")

	for _, wonder := range r.db.Wonders {
		if wonder.Name == name {
			return c.JSON(http.StatusOK, wonder)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Could not find wonder")
}
