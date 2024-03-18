package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonitienda/kadai/backend-golang-rest/pkg/tasks"
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

func TestGetAddDeleteTasksUser(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"

	newTaskBody := []byte(`{ "title": "Do something", "description": "some details about the task" }`)

	router := setupRouter()

	w := httptest.NewRecorder()
	getTasksRequest, _ := http.NewRequest("GET", "/tasks", nil)
	getTasksRequest.Header.Add("X-User-Id", userId)

	addTasksRequest, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(newTaskBody))
	addTasksRequest.Header.Add("X-User-Id", userId)

	router.ServeHTTP(w, getTasksRequest)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())

	router.ServeHTTP(w, addTasksRequest)

	assert.Equal(t, 200, w.Code)

	fmt.Println("New Task Response:", w.Body.String())

	var newTaskResponse tasks.SirenResponse[tasks.TaskProperties]
	err := json.Unmarshal(w.Body.Bytes(), &newTaskResponse)

	assert.NoError(t, err)

	fmt.Println("New Task Response:", newTaskResponse)

	assert.Equal(t, "Do something", newTaskResponse.Properties.Title)
	assert.Equal(t, "some details about the task", newTaskResponse.Properties.Description)
	assert.Equal(t, "pending", newTaskResponse.Properties.Status)
	assert.Equal(t, userId, newTaskResponse.Properties.OwnerID)

	// TODO - Assert that the format is a valid UUID
	assert.NotEmpty(t, newTaskResponse.Properties.ID)

}
