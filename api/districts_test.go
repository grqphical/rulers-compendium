package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grqphical07/rulers-compendium/database"
	_ "github.com/grqphical07/rulers-compendium/test_init"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetDistricts(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/districts/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	if assert.NoError(t, router.GetDistricts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_districts := make([]database.District, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_districts)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Districts, recieved_districts)
	}

	e = echo.New()
	req = httptest.NewRequest(http.MethodGet, "/api/v1/districts?limit=2", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	router = NewRouter(&db, e)

	if assert.NoError(t, router.GetDistricts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_districts := make([]database.District, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_districts)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, []database.District{db.Districts[0], db.Districts[1]}, recieved_districts)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/districts?limit=-1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Error(t, router.GetDistricts(c))
}

func TestGetDistrict(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/districts/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("City Center")

	if assert.NoError(t, router.GetDistrict(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_districts := database.District{}
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_districts)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Districts[0], recieved_districts)
	}

	e = echo.New()
	req = httptest.NewRequest(http.MethodGet, "/api/v1/leaders/", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	router = NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("NotADistrict")

	assert.Error(t, router.GetDistrict(c))
}
