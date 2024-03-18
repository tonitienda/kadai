package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTasksAnonymousUser(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestGetTasksNewUser(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Add("X-User-Id", "d5ba1349-7954-4efd-9913-1a68dfd008d8")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())

}
