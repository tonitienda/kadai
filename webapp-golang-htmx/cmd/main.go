package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"

	authenticator "github.com/tonitienda/kadai/webapp-golang-htmx/pkg/auth"
	"github.com/tonitienda/kadai/webapp-golang-htmx/pkg/tasks"
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

		http.Redirect(w, r, auth.AuthCodeURL(state), http.StatusTemporaryRedirect)

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
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx, code)
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

	http.HandleFunc("/task-list", func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("id_token")

		if err != nil {
			http.Error(w, "No token found", http.StatusUnauthorized)
			return
		}

		// TODO - DELETE THIS!!
		fmt.Println("Token: ", token.Value)

		templates.ExecuteTemplate(w, "task-list.html", nil)

	})

	http.ListenAndServe(":3000", nil)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
