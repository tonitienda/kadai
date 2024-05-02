package main

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/tonitienda/kadai/bdd-go/pkg/api"
)

type TestsImpl interface {
	UserNotLoggedIn(user string) error
	UserIsLoggedIn(user string) error
	UserRequestsListOfTasks(user string) error
	UserAddsTask(user string, title string, description string) error
	UserShouldHaveNTasks(user string, numTasks int) error
	TheListOfTasksShouldBeEmpty() error
	UnauthorizedErrorReturned() error
	BadRequestReturned() error
	SuccessfulRequestReturned() error
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	InitializeScenario2(ctx, api.NewApiTests())
}

func InitializeScenario2(ctx *godog.ScenarioContext, impl TestsImpl) {

	ctx.Given(`^(.*) is not logged in$`, impl.UserNotLoggedIn)
	ctx.Given(`^(.*) is logged in$`, impl.UserIsLoggedIn)
	ctx.When(`^(.*) requests the list of tasks$`, impl.UserRequestsListOfTasks)
	ctx.Step(`^(.*) adds a task with title "([^"]*)" and description "([^"]*)"$`, impl.UserAddsTask)
	ctx.Then(`^the list of tasks should be empty$`, impl.TheListOfTasksShouldBeEmpty)
	ctx.Then(`^(.*) should have (\d+) tasks?$`, impl.UserShouldHaveNTasks)
	ctx.Then(`^the request should be refused because the user is unauthorized$`, impl.UnauthorizedErrorReturned)
	ctx.Then(`^the request should be successful$`, impl.SuccessfulRequestReturned)
	ctx.Then(`^the request should be refused because the data is not valid$`, impl.BadRequestReturned)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
