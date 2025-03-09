package clickup

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetWorkspaces() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	token := os.Getenv("CLICKUP_API_TOKEN")

	url := "https://api.clickup.com/api/v2/team"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	defer resp.Body.Close()
}
