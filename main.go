package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
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

func createRepo(cmd *cobra.Command, args []string) {
	var repoName string

	token := "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"

	prompt := &survey.Input{
		Message: "Enter repo name: ",
	}

	survey.AskOne(prompt, &repoName, nil)

	var clone bool

	var isPrivate bool
	promptt := &survey.Confirm{
		Message: "Is is private?: ",
	}

	survey.AskOne(promptt, &isPrivate, nil)

	promptt = &survey.Confirm{
		Message: "Do you want to clone it?: ",
	}

	survey.AskOne(promptt, &clone, nil)
	repo := CreateRepoRequest{
		Name:        repoName,
		Description: "A new repo created using github api",
		Private:     isPrivate,
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

		if clone {
			err := gitClone(createRepoResponse.URL)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Repo cloned successfully")
		}
	} else {
		fmt.Println("Error creating repo")
		fmt.Println(string(respBody))
	}
}

func deleteRepo() {

}

func gitClone(url string) error {
	cmd := exec.Command("git", "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro cloning repo: %w", err)
	}

	return nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "github-create-repo",
		Short: "Create a new repo on Github",
		Args:  cobra.NoArgs,
		Run:   createRepo,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
