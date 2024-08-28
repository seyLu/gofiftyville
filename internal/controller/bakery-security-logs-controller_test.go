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

func TestGetBakerySecurityLogs(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetBakerySecurityLogsBadDate(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	q := req.URL.Query()
	q.Add("date", "January 20, 202")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetBakerySecurityLogsGoodDate(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	q := req.URL.Query()
	q.Add("date", "January 20, 2021")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetBakerySecurityLogsBadTime(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	q := req.URL.Query()
	q.Add("time", "11:00 P")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetBakerySecurityLogsGoodTime(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	q := req.URL.Query()
	q.Add("time", "11:00 PM")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetBakerySecurityLogsBadTime2(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	q := req.URL.Query()
	q.Add("time-2", "11:00 P")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetBakerySecurityLogsGoodTime2(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bakery-security-logs", nil)
	q := req.URL.Query()
	q.Add("time-2", "11:00 PM")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
