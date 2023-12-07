package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllSpaces(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/spaces/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	w = httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+getAuthToken())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
