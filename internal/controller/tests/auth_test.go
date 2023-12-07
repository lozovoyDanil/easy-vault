package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	w := httptest.NewRecorder()
	query := []byte(`{"name": "test","email": "testing.com", "password": "test"}`)
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(query))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "{\"id\":3}", w.Body.String())
}

func TestSignIn(t *testing.T) {
	w := httptest.NewRecorder()
	query := []byte(`{"login": "testing.com", "password": "test"}`)
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(query))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
