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

func TestGetCivilizations(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/civilizations", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	if assert.NoError(t, router.GetCivilizations(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_civilizations := make([]database.Civilization, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_civilizations)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Civilizations, recieved_civilizations)
	}

	// Test Limit Query Param
	req = httptest.NewRequest(http.MethodGet, "/api/v1/civilizations?limit=1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, router.GetCivilizations(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_civilization := make([]database.Civilization, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_civilization)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, []database.Civilization{db.Civilizations[0]}, recieved_civilization)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/civilizations?limit=-1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Error(t, router.GetCivilizations(c))
}

func TestGetCivilization(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/civilizations/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("American")

	if assert.NoError(t, router.GetCivilization(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_civilization := database.Civilization{}
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_civilization)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Civilizations[0], recieved_civilization)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/civilizations/", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	router = NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("NotACiv")

	assert.Error(t, router.GetCivilization(c))
}
