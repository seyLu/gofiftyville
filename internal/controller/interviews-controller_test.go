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

func TestGetInterviews(t *testing.T) {
	store.SetupTestDB()
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/interviews", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
