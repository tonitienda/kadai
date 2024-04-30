package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Undo struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

func GetTasks(token string) ([]Task, error) {
	fmt.Println("Getting tasks...")

	req, err := http.NewRequest(http.MethodGet, "http://backend:8080/v0/tasks", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		// TODO - See this
		return []Task{}, nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		// TODO - See this
		return []Task{}, nil
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		fmt.Printf("client: got non-200 status code: %d\n", res.StatusCode)
		// TODO - See this
		return []Task{}, nil
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		// TODO - See this
		return []Task{}, nil
	}
	fmt.Printf("Get Tasks: response body: %s\n", resBody)

	// Deserialize json to an array of Tasks
	tasks := []Task{}

	err = json.Unmarshal(resBody, &tasks)

	if err != nil {
		fmt.Printf("client: could not deserialize json: %s\n", err)
		return nil, err
	}

	return tasks, nil

}

func AddTask(token string, title string, description string) error {
	fmt.Println("Adding task...")

	req, err := http.NewRequest(http.MethodPost, "http://backend:8080/v0/tasks", bytes.NewReader([]byte(
		fmt.Sprintf(`{"title": "%s", "description": "%s"}`, title, description),
	)))

	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		// TODO - See this
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		// TODO - See this
		return err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		fmt.Printf("client: got non-200 status code: %d\n", res.StatusCode)
		// TODO - See this
		return fmt.Errorf("non-200 status code: %d", res.StatusCode)
	}

	return nil

}

func UndoTaskChange(token string, url string, method string) error {
	fmt.Println("Undoing task change...")

	req, err := http.NewRequest(method, "http://backend:8080"+url, nil)

	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		// TODO - See this
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		// TODO - See this
		return err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusAccepted {
		fmt.Printf("client: got non-202 status code: %d\n", res.StatusCode)
		// TODO - See this
		return fmt.Errorf("non-202 status code: %d", res.StatusCode)
	}

	return nil
}

func DeleteTask(token string, taskID string) (Undo, error) {
	fmt.Println("Deleting task...")

	req, err := http.NewRequest(http.MethodDelete, "http://backend:8080/v0/tasks/"+taskID, nil)

	undo := Undo{}

	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		// TODO - See this
		return undo, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		// TODO - See this
		return undo, err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusAccepted {
		fmt.Printf("client: got non-202 status code: %d\n", res.StatusCode)
		// TODO - See this
		return undo, fmt.Errorf("non-202 status code: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		// TODO - See this
		return undo, nil
	}
	fmt.Printf("Delete Task: response body: %s\n", resBody)

	// Deserialize json to an array of Tasks

	err = json.Unmarshal(resBody, &undo)

	if err != nil {
		fmt.Printf("client: could not deserialize json: %s\n", err)
		return undo, err
	}

	return undo, nil

}
