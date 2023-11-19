package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grqphical07/rulers-compendium/database"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetImprovements(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/improvements/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	if assert.NoError(t, router.GetImprovements(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_improvements := make([]database.Improvement, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_improvements)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, db.Improvements, recieved_improvements)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/improvements?limit=1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, router.GetImprovements(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_improvements := make([]database.Improvement, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_improvements)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, []database.Improvement{db.Improvements[0]}, recieved_improvements)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/improvements?limit=-1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Error(t, router.GetImprovements(c))
}

func TestGetImprovement(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/improvements", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("Farm")

	if assert.NoError(t, router.GetImprovement(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_improvement := database.Improvement{}
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_improvement)
		if err != nil {
			panic(err)
		}
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/improvements", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("NotAnImprovement")

	assert.Error(t, router.GetImprovement(c))
}
