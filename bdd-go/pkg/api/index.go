package api

import (
	"fmt"
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
		fmt.Println("correct")
	}
	return nil
}
func (a *ApiTests) TheListOfTasksShouldBeEmpty() error {
	if len(a.LastRequestedTasks) != 0 {
		return fmt.Errorf("the list of tasks should be empty")
	}
	return nil
}
func (a *ApiTests) UnauthorizedErrorReturned() error {
	if a.LastRequestStatus != http.StatusUnauthorized {
		return fmt.Errorf("expected status %d but %d was returned", http.StatusUnauthorized, a.LastRequestStatus)
	}

	return nil
}

func (a *ApiTests) SuccessfulRequestReturned() error {
	if a.LastRequestStatus != http.StatusOK {
		return fmt.Errorf("expected status %d but %d was returned", http.StatusOK, a.LastRequestStatus)
	}

	return nil
}
