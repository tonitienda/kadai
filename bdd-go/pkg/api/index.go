package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ApiTests struct {
	UserTokens         map[string]string
	UserTasks          map[string][]Task
	LastRequestedTasks []Task
	LastRequestStatus  int
	BackendUrl         string
}

func NewApiTests() *ApiTests {
	return &ApiTests{
		UserTokens:         map[string]string{},
		LastRequestedTasks: []Task{},
		BackendUrl:         os.Getenv("BACKEND_BASE_URL"),
	}
}

func (a *ApiTests) UserNotLoggedIn(user string) error {
	delete(a.UserTokens, user)
	return nil
}

func (a *ApiTests) UserIsLoggedIn(user string) error {
	fmt.Println(user, " logged in")
	token, err := userLogin(user)

	if err != nil {
		return err
	}

	a.UserTokens[user] = token
	return nil
}

func (a *ApiTests) UserRequestsListOfTasks(user string) error {
	url := fmt.Sprintf("%s/v0/tasks", a.BackendUrl)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	token, isUserLoggedIn := a.UserTokens[user]

	fmt.Printf("User tokens: %v", a.UserTokens)
	if isUserLoggedIn {
		fmt.Println("User", user, "is logged in")
		req.Header.Add("Authorization", "Bearer "+token)
	}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	a.LastRequestStatus = res.StatusCode

	if res.StatusCode == http.StatusOK {
		// Deserialize tasks
		fmt.Println("Deserializing tasks...")

		body, err := io.ReadAll(res.Body)
		fmt.Println("Response body:", string(body))
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, &a.LastRequestedTasks); err != nil {
			return err
		}

		fmt.Printf("LastRequestedTasks: %v\n", a.LastRequestedTasks)
	}
	return nil
}

func (a *ApiTests) UserAddsTask(user string, title string, description string) error {
	url := fmt.Sprintf("%s/v0/tasks", a.BackendUrl)

	var jsonStr = []byte(fmt.Sprintf(`{"title":"%s","description":"%s"}`, title, description))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	token, isUserLoggedIn := a.UserTokens[user]
	req.Header.Add("content-type", "application/json")

	fmt.Printf("User tokens: %v", a.UserTokens)
	if isUserLoggedIn {
		fmt.Println("User", user, "is logged in")
		req.Header.Add("Authorization", "Bearer "+token)
	}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	a.LastRequestStatus = res.StatusCode

	return nil
}

func (a *ApiTests) TheListOfTasksShouldBeEmpty() error {
	if len(a.LastRequestedTasks) != 0 {
		return fmt.Errorf("the list of tasks should be empty")
	}
	return nil
}

func (a *ApiTests) UserShouldHaveNTasks(user string, numTasks int) error {
	a.UserRequestsListOfTasks(user)

	if len(a.LastRequestedTasks) != numTasks {
		return fmt.Errorf("the list of tasks should have %d items but %d were found", numTasks, len(a.LastRequestedTasks))
	}
	return nil
}

func (a *ApiTests) UnauthorizedErrorReturned() error {
	if a.LastRequestStatus != http.StatusUnauthorized {
		return fmt.Errorf("expected status %d but %d was returned", http.StatusUnauthorized, a.LastRequestStatus)
	}

	return nil
}

func (a *ApiTests) BadRequestReturned() error {
	if a.LastRequestStatus != http.StatusBadRequest {
		return fmt.Errorf("expected status %d but %d was returned", http.StatusBadRequest, a.LastRequestStatus)
	}

	return nil
}

func (a *ApiTests) SuccessfulRequestReturned() error {
	if a.LastRequestStatus != http.StatusOK {
		return fmt.Errorf("expected status %d but %d was returned", http.StatusOK, a.LastRequestStatus)
	}

	return nil
}
