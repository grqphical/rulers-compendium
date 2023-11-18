package api

import (
	"example/civ6-api/database"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func getLeadersByCivilization(civilzation string, limit int, db *database.Database) []database.Leader {
	civ_leaders := make([]database.Leader, 0)

	for _, leader := range db.Leaders {
		if limit != 0 && len(civ_leaders) == limit {
			return civ_leaders
		}

		if strings.ToLower(leader.Civilization) == strings.ToLower(civilzation) {
			civ_leaders = append(civ_leaders, leader)
		}
	}

	return civ_leaders
}

func (r *Router) Index(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}

func (r *Router) GetLeaders(c echo.Context) error {
	limit := c.QueryParam("limit")
	civilzation := c.QueryParam("civilization")

	if limit == "" {
		if civilzation != "" {
			return c.JSON(http.StatusOK, getLeadersByCivilization(civilzation, 0, r.db))
		}
		leaders := r.db.Leaders
		return c.JSON(http.StatusOK, leaders)
	} else {
		limit, err := strconv.Atoi(limit)
		if err != nil || limit <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit value. Must be number above zero")
		}

		if civilzation != "" {
			return c.JSON(http.StatusOK, getLeadersByCivilization(civilzation, limit, r.db))
		}

		leaders := r.db.Leaders[:limit]
		return c.JSON(http.StatusOK, leaders)
	}

}

func (r *Router) GetLeader(c echo.Context) error {
	name := c.Param("name")
	name = strings.ReplaceAll(name, "%20", " ")

	for _, leader := range r.db.Leaders {
		if leader.Name == name {
			return c.JSON(http.StatusOK, leader)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Could not find leader")
}
