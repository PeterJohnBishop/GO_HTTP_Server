package clickup

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetAccessToken(client_id string, client_secret string, code string) ([]byte, error) {
	url := "https://api.clickup.com/api/v2/oauth/token"

	requestBody, _ := json.Marshal(map[string]string{
		"client_id":     client_id,
		"client_secret": client_secret,
		"code":          code,
	})

	req, err := http.NewRequest("GET", url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetAuthorizedUser() ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("CLICKUP_PK")

	url := "https://api.clickup.com/api/v2/user"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetWorkspaces() ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("CLICKUP_PK")

	url := "https://api.clickup.com/api/v2/team"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}
