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

func TestNoFinalAnswer(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/final-answer", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestWrongFinalAnswer(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/final-answer", nil)
	q := req.URL.Query()
	q.Add("thief", "John")
	q.Add("city", "Foo City")
	q.Add("accomplice", "Doe")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCorrectFinalAnswer(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/final-answer", nil)
	q := req.URL.Query()
	q.Add("thief", "Bruce")
	q.Add("city", "New York City")
	q.Add("accomplice", "Robin")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
