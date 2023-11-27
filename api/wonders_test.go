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

func TestGetWonders(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/wonders", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	if assert.NoError(t, router.GetWonders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_wonders := make([]database.Wonder, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_wonders)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Wonders, recieved_wonders)
	}

	// Test Limit Query Param
	req = httptest.NewRequest(http.MethodGet, "/api/v1/wonders?limit=1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, router.GetWonders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_wonder := make([]database.Wonder, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_wonder)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, []database.Wonder{db.Wonders[0]}, recieved_wonder)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/wonders?limit=-1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Error(t, router.GetWonders(c))
}

func TestGetWonder(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/wonders", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("Alhambra")

	if assert.NoError(t, router.GetWonder(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_wonder := database.Wonder{}
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_wonder)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Wonders[0], recieved_wonder)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/wonders/", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	router = NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("NotAWonder")

	assert.Error(t, router.GetCivilization(c))
}
