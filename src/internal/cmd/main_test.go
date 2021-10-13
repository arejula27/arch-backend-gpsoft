package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestCreateBook(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	postBody, _ := json.Marshal(map[string]string{
		"id":   "1",
		"name": "quijote",
	})
	responseBody := bytes.NewBuffer(postBody)
	println(responseBody)
	println()
	println()

	req, _ := http.NewRequest("POST", "/save", responseBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}
