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

func TestGetLeaders(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/leaders", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	if assert.NoError(t, router.GetLeaders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_leaders := make([]database.Leader, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_leaders)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Leaders, recieved_leaders)
	}

	// Test Limit Query Param
	req = httptest.NewRequest(http.MethodGet, "/api/v1/leaders?limit=1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, router.GetLeaders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_leader := make([]database.Leader, 0)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_leader)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, []database.Leader{db.Leaders[0]}, recieved_leader)
	}

	// Test Civilization Param
	req = httptest.NewRequest(http.MethodGet, "/api/v1/leaders?civilization=Macedonia", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, router.GetLeaders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_leader := make([]database.Leader, 1)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_leader)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, []database.Leader{db.Leaders[1]}, recieved_leader)
	}

	// Test Civilization and Limit Param
	req = httptest.NewRequest(http.MethodGet, "/api/v1/leaders?civilization=America&limit=1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, router.GetLeaders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_leader := make([]database.Leader, 1)
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_leader)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, []database.Leader{db.Leaders[0]}, recieved_leader)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/leaders?limit=-1", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Error(t, router.GetLeaders(c))
}

func TestGetLeader(t *testing.T) {
	db := database.ReadDatabase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/leaders/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	router := NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("Alexander")

	if assert.NoError(t, router.GetLeader(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recieved_leader := database.Leader{}
		err := json.Unmarshal(rec.Body.Bytes(), &recieved_leader)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, db.Leaders[1], recieved_leader)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/leaders/", nil)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	router = NewRouter(&db, e)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("NotALeader")

	assert.Error(t, router.GetLeader(c))
}
