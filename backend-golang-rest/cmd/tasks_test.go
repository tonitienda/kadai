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
	req, _ := http.NewRequest("GET", "/v0/tasks", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestGetTasksNewUser(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v0/tasks", nil)
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

func getUserTasks(t *testing.T, router *gin.Engine, userId string) []tasks.TaskResponse {
	getTasksRequest, _ := http.NewRequest("GET", "/v0/tasks", nil)
	getTasksRequest.Header.Add("X-User-Id", userId)
	w := sendRequest(router, getTasksRequest)

	var tasks []tasks.TaskResponse
	err := json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.NoError(t, err)
	return tasks
}

func addTask(t *testing.T, router *gin.Engine, userId string, title string, description string) string {
	newTaskBody := []byte(`{ "title": "` + title + `", "description": "` + description + `" }`)

	addTasksRequest, _ := http.NewRequest("POST", "/v0/tasks", bytes.NewBuffer(newTaskBody))
	addTasksRequest.Header.Add("X-User-Id", userId)

	w := sendRequest(router, addTasksRequest)

	assert.Equal(t, 200, w.Code)

	var newTaskResponse tasks.TaskResponse
	err := json.Unmarshal(w.Body.Bytes(), &newTaskResponse)

	assert.NoError(t, err)

	return newTaskResponse.ID
}

func deleteTask(t *testing.T, router *gin.Engine, userId string, taskId string) {
	deleteTaskRequest, _ := http.NewRequest("DELETE", "/v0/tasks/"+taskId, nil)
	deleteTaskRequest.Header.Add("X-User-Id", userId)

	w := sendRequest(router, deleteTaskRequest)

	assert.Equal(t, 202, w.Code)
}

func TestGetAddDeleteTasksUser(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"

	router := setupRouter()
	tasks := getUserTasks(t, router, userId)
	assert.Empty(t, tasks)

	newTaskID := addTask(t, router, userId, "Do something", "some details about the task")

	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)
	assert.Equal(t, newTaskID, tasks[0].ID)

	deleteTask(t, router, userId, newTaskID)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 0)
}
