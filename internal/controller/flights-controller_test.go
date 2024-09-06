package controller_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seyLu/gofiftyville/internal/router"
	"github.com/seyLu/gofiftyville/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestGetFlights(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/flights", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetFlightsBadDate(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/flights", nil)
	q := req.URL.Query()
	q.Add("date", "January 20, 202")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetFlightsGoodDate(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/flights", nil)
	q := req.URL.Query()
	q.Add("date", "January 20, 2021")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetFlightsBadPassportNumbers(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/flights", nil)
	q := req.URL.Query()
	q.Add("passport-number", "4041498045")
	q.Add("passport-number", "qewqedasd")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetFlightsGoodPassportNumbers(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/flights", nil)
	q := req.URL.Query()
	q.Add("passport-number", "4041498045")
	q.Add("passport-number", "7441135547")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
