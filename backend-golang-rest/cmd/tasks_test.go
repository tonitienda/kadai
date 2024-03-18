package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func sendRequest(router *gin.Engine, request *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w
}

func TestGetAddDeleteTasksUser(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"

	newTaskBody := []byte(`{ "title": "Do something", "description": "some details about the task" }`)

	router := setupRouter()

	getTasksRequest, _ := http.NewRequest("GET", "/tasks", nil)
	getTasksRequest.Header.Add("X-User-Id", userId)

	addTasksRequest, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(newTaskBody))
	addTasksRequest.Header.Add("X-User-Id", userId)

	w := sendRequest(router, getTasksRequest)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())

	w = sendRequest(router, addTasksRequest)

	assert.Equal(t, 200, w.Code)

	var newTaskResponse tasks.SirenResponse[tasks.TaskProperties]
	err := json.Unmarshal(w.Body.Bytes(), &newTaskResponse)

	assert.NoError(t, err)

	assert.Equal(t, "Do something", newTaskResponse.Properties.Title)
	assert.Equal(t, "some details about the task", newTaskResponse.Properties.Description)
	assert.Equal(t, "pending", newTaskResponse.Properties.Status)
	assert.Equal(t, userId, newTaskResponse.Properties.OwnerID)

	// TODO - Assert that the format is a valid UUID
	assert.NotEmpty(t, newTaskResponse.Properties.ID)

	// Delete task url
	var deleteTaskURL string
	for _, action := range newTaskResponse.Actions {
		if action.Name == "delete-task" {
			deleteTaskURL = action.Href
			break
		}
	}

	assert.NotEmpty(t, deleteTaskURL)

	deleteTaskRequest, _ := http.NewRequest("DELETE", deleteTaskURL, nil)
	deleteTaskRequest.Header.Add("X-User-Id", userId)

	w = sendRequest(router, deleteTaskRequest)

	assert.Equal(t, http.StatusAccepted, w.Code)

	// Assert that the task list is empty again
	w = sendRequest(router, getTasksRequest)

	assert.Equal(t, "[]", w.Body.String())

}
