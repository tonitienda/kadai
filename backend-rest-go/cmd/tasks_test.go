package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tonitienda/kadai/backend-rest-go/pkg/tasks"
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

func deleteTask(t *testing.T, router *gin.Engine, userId string, taskId string, status int) {
	deleteTaskRequest, _ := http.NewRequest("DELETE", "/v0/tasks/"+taskId, nil)
	deleteTaskRequest.Header.Add("X-User-Id", userId)

	w := sendRequest(router, deleteTaskRequest)

	assert.Equal(t, status, w.Code)
}

func deleteTaskAccepted(t *testing.T, router *gin.Engine, userId string, taskId string) {
	deleteTask(t, router, userId, taskId, http.StatusAccepted)
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

	deleteTaskAccepted(t, router, userId, newTaskID)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 0)
}

func TestGetTasksNoAuthHeader(t *testing.T) {
	router := setupRouter()

	getTasksRequest, _ := http.NewRequest("GET", "/v0/tasks", nil)
	w := sendRequest(router, getTasksRequest)
	assert.Equal(t, 401, w.Code)
}

func TestGetTasksNoUser(t *testing.T) {
	router := setupRouter()

	getTasksRequest, _ := http.NewRequest("GET", "/v0/tasks", nil)
	getTasksRequest.Header.Add("X-User-Id", "")
	w := sendRequest(router, getTasksRequest)
	assert.Equal(t, 401, w.Code)
}

func TestTryDeleteDifferentUserTask(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"
	otherUserId := "d001ab25-83ee-4625-81f8-f80d06e2232f"

	router := setupRouter()

	newTaskID := addTask(t, router, userId, "Do something", "some details about the task")

	tasks := getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, otherUserId, newTaskID, http.StatusForbidden)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, userId, newTaskID, http.StatusAccepted)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 0)

}

func TestTryDeleteWithoutTaskID(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"
	router := setupRouter()

	newTaskID := addTask(t, router, userId, "Do something", "some details about the task")

	tasks := getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	// This should be a 400 since task ID is missing but
	// Gin finds that the param is not correct and returns a 404
	// Because there is no endpoint /tasks with DELETE verb
	deleteTask(t, router, userId, "", http.StatusNotFound)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, userId, newTaskID, http.StatusAccepted)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 0)

}
func TestTryDeleteWithNonExistingTaskID(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"
	router := setupRouter()

	newTaskID := addTask(t, router, userId, "Do something", "some details about the task")

	tasks := getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, userId, "7227f1ac-b965-4716-9d2d-ed4c19acf585", http.StatusNotFound)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, userId, newTaskID, http.StatusAccepted)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 0)

}

func TestTryDeleteWithWrongTaskID(t *testing.T) {
	userId := "3a34557c-d1bd-4286-b639-e1c915209a12"
	router := setupRouter()

	newTaskID := addTask(t, router, userId, "Do something", "some details about the task")

	tasks := getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, userId, "-", http.StatusBadRequest)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 1)

	deleteTask(t, router, userId, newTaskID, http.StatusAccepted)
	tasks = getUserTasks(t, router, userId)
	assert.Len(t, tasks, 0)

}
