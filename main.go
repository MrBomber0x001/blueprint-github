package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

type CreateRepoResponse struct {
	Name   string `json:"name"`
	URL    string `json:"html_url"`
	Errors []struct {
		Message string `json:"message`
	} `json:"errors"`
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "github-create-repo",
		Short: "Create a new repo on Github",
		Args:  cobra.ExactArgs(1),
		Run:   createRepo,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createRepo(cmd *cobra.Command, args []string) {
	repoName := args[0]

	token := "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"

	repo := CreateRepoRequest{
		Name:        repoName,
		Description: "A new repo created using github api",
		Private:     true,
	}
	body, _ := json.Marshal(repo)

	url := "https://api.github.com/user/repos"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating repo", err)
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusCreated {
		var createRepoResponse CreateRepoResponse
		json.Unmarshal(respBody, &createRepoResponse)

		fmt.Println("repo created successfully")

		fmt.Println("name", createRepoResponse.Name)
		fmt.Println("url:", createRepoResponse.URL)

	} else {
		fmt.Println("Error creating repo")
		fmt.Println(string(respBody))
	}
}

func deleteRepo() {

}
