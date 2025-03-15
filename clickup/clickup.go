package clickup

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func saveToEnv(key, value string) error {
	envFile := ".env"

	envMap, err := godotenv.Read(envFile)
	if err != nil {
		return err
	}

	envMap[key] = value

	err = godotenv.Write(envMap, envFile)
	if err != nil {
		return err
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//log.Printf("Saved %s=%s to .env\n", key, value)
	return nil
}

func GetAccessToken(client_id string, client_secret string) ([]byte, error) {
	os.Clearenv()
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	code := os.Getenv("AUTH_CODE")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/oauth/token?client_id=%s&client_secret=%s&code=%s", client_id, client_secret, code)

	req, err := http.NewRequest("POST", url, nil)
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
	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return nil, err
	}
	err = saveToEnv("OAUTH_TOKEN", authResp.AccessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

func GetAuthorizedUser() ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := "https://api.clickup.com/api/v2/user"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
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
	token := os.Getenv("OAUTH_TOKEN")

	url := "https://api.clickup.com/api/v2/team"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetSpaces(team_id string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/team/%s/space", team_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetFolders(space_id string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/space/%s/folder", space_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetFolderlessLists(space_id string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/space/%s/list", space_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetLists(folder_id string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/folder/%s/list", folder_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetTasks(list_id string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/list/%s/task", list_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func GetTask(task_id string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("OAUTH_TOKEN")

	url := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s", task_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}
