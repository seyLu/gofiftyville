package controller_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seyLu/gofiftyville/internal/router"
	"github.com/seyLu/gofiftyville/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestGetPhoneCalls(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPhoneCallsBadDate(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("date", "January 20, 202")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetPhoneCallsGoodDate(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("date", "January 20, 2021")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPhoneCallsBadDurationInequality(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("duration-inequality", "=")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetPhoneCallsGoodDurationInequality(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("duration-inequality", ">")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPhoneCallsBadDuration(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("duration", "dsadasda")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}

func TestGetPhoneCallsGoodDuration(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("duration", "100")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)
	fmt.Println(w)

	assert.Equal(t, 200, w.Code)
}

func TestGetPhoneCallsGoodCallers(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/phone-calls", nil)
	q := req.URL.Query()
	q.Add("caller", "(123) 555-5144")
	q.Add("caller", "(741) 555-1748")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)
	fmt.Println(w)

	assert.Equal(t, 200, w.Code)
}
