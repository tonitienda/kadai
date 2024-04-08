package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func GetTasks(token string) ([]Task, error) {
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
	fmt.Printf("client: response body: %s\n", resBody)

	// Deserialize json to an array of Tasks
	tasks := []Task{}

	err = json.Unmarshal(resBody, &tasks)

	if err != nil {
		fmt.Printf("client: could not deserialize json: %s\n", err)
		return nil, err
	}

	return tasks, nil

}
