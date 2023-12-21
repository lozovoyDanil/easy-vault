package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	env, err := InitEnv()
	if err != nil {
		panic(err)
	}
	defer env.DB.Close()

	router = env.Handler.InitRoutes()

	exitCode := m.Run()

	// err = env.Remove()
	// if err != nil {
	// 	panic(err)
	// }

	os.Exit(exitCode)
}

func getAuthToken() string {
	w := httptest.NewRecorder()
	query := []byte(`{"login": "danilozovoy@gmail.com", "password": "12345678"}`)
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(query))
	router.ServeHTTP(w, req)

	fullToken := w.Body.String()
	token := strings.Split(fullToken, ":")[1]

	return token[1 : len(token)-2]
}
