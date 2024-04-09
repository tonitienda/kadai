package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	authenticator "github.com/tonitienda/kadai/webapp-golang-htmx/pkg/auth"
	"github.com/tonitienda/kadai/webapp-golang-htmx/pkg/tasks"
	"golang.org/x/oauth2"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func init() {

}

// TODO - Make cookie secure and httponly
func newCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: value,
		//HttpOnly: true,
		//SameSite: http.SameSiteStrictMode,
		//Secure:   true,
	}
}

func main() {

	auth, err := authenticator.New()

	if err != nil {
		fmt.Println("Error initializing authenticator: ", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := []tasks.Task{}

		cookie, _ := r.Cookie("access_token")
		if cookie != nil {
			t, err = tasks.GetTasks(cookie.Value)

			if err != nil {
				fmt.Println("Failed to add task: ", err)
				http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
				return
			}
		}

		templates.ExecuteTemplate(w, "index.html", struct{ Tasks []tasks.Task }{
			Tasks: t,
		})
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		state, err := generateRandomState()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		http.SetCookie(w, newCookie("auth_state", state))

		authUrl := auth.AuthCodeURL(state, oauth2.SetAuthURLParam("audience", os.Getenv("AUTH0_AUDIENCE")))
		fmt.Println("Redirecting to: ", authUrl)

		http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)

	})

	http.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("POST /tasks")
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		t := []tasks.Task{}

		// Retrieve the title and description from the form
		title := r.FormValue("title")
		description := r.FormValue("description")

		cookie, _ := r.Cookie("access_token")
		if cookie != nil {
			err := tasks.AddTask(cookie.Value, title, description)

			if err != nil {
				fmt.Println("Failed to add task: ", err)
				http.Error(w, "Failed to add task", http.StatusInternalServerError)
				return
			}

			t, err = tasks.GetTasks(cookie.Value)

			if err != nil {
				fmt.Println("Failed to get tasks: ", err)
				http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
				return
			}

			fmt.Println("Tasks: ", t)
		}

		templates.ExecuteTemplate(w, "task-list.html", struct{ Tasks []tasks.Task }{
			Tasks: t,
		})
	})

	http.HandleFunc("DELETE /tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("DELETE /tasks")

		path := strings.TrimPrefix(r.URL.Path, "/tasks/")
		id := strings.TrimSuffix(path, "/")
		fmt.Fprintf(w, "Deleting task with ID: %s", id)

		t := []tasks.Task{}

		cookie, _ := r.Cookie("access_token")
		if cookie != nil {
			err := tasks.DeleteTask(cookie.Value, id)

			if err != nil {
				fmt.Println("Failed to delete task: ", err)
				http.Error(w, "Failed to delete task", http.StatusInternalServerError)
				return
			}

			t, err = tasks.GetTasks(cookie.Value)

			if err != nil {
				fmt.Println("Failed to get tasks: ", err)
				http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
				return
			}

			fmt.Println("Tasks: ", t)
		}

		templates.ExecuteTemplate(w, "task-list.html", struct{ Tasks []tasks.Task }{
			Tasks: t,
		})
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		storedState, err := r.Cookie("auth_state")
		if err != nil {
			// Handle cookie not found error
			http.Error(w, "Auth token cookie not found", http.StatusUnauthorized)
			return
		}

		if r.URL.Query().Get("state") != storedState.Value {
			http.Error(w, "Invalid state parameter.", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")
		ctx := context.Background()

		// TODO - See what to do here, when the code is not passed to the callback
		if code == "" {
			http.Error(w, "Code not found in request.", http.StatusBadRequest)
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx, code, oauth2.SetAuthURLParam("audience", os.Getenv("AUTH0_AUDIENCE")))
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Failed to exchange an authorization code for a token."))
			return
		}

		idToken, err := auth.VerifyIDToken(ctx, token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to verify ID Token."))
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// TODO - Serialize this as JSON and store it as one item.
		http.SetCookie(w, newCookie("access_token", token.AccessToken))
		http.SetCookie(w, newCookie("id_token", token.Extra("id_token").(string)))
		http.SetCookie(w, newCookie("user_nickname", profile["nickname"].(string)))
		http.SetCookie(w, newCookie("user_picture", profile["picture"].(string)))

		fmt.Println("Callback executed, redirecting to /")
		http.Redirect(w, r, "http://localhost:3000", http.StatusFound)

	})

	http.ListenAndServe(":3000", nil)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.RawURLEncoding.EncodeToString(b)

	return state, nil
}
