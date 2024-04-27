package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func userLogin(user string) (string, error) {

	url := os.Getenv("AUTH0_ISSUER_BASE_URL") + "/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"" + os.Getenv("AUTH0_KADAI_BACKEND_CLIENT_ID") + "\",\"client_secret\":\"" + os.Getenv("AUTH0_KADAI_BACKEND_CLIENT_SECRET") + "\",\"audience\":\"" + os.Getenv("AUTH0_AUDIENCE") + "\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	fmt.Printf("OK: data=%#v\n", data)

	token, ok := data["access_token"]

	if !ok {
		return "", fmt.Errorf("the token could not be obtained")
	}

	return fmt.Sprint(token), nil
}
