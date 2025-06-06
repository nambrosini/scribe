package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/nambrosini/scribe/internal/config"
)

const (
	// URL        = "https://api.mistral.ai/v1/chat/completions"
	URL        = "http://localhost:11434/api/chat"
	TICKET_KEY = "TICKET"
)

type Model string

const (
	MistralLargeLatest Model = "mistral-large-latest"
	MistrallLatest     Model = "mistral:latest"
	MistralSmallLatest Model = "mistral-small:latest"
)

type Role string

const (
	System Role = "system"
	User   Role = "user"
)

func SendRequest(cfg config.AppConfig) (string, error) {
	messages, err := BuildMessages(cfg.Commit.Template, cfg.Commit.Issue, cfg.Commit.Type)
	if err != nil {
		return "", nil
	}
	requestBody := RequestBody{
		Model:    Model(cfg.Model.Name),
		Messages: messages,
		Stream:   false,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", nil
	}

	req, err := http.NewRequest(http.MethodPost, cfg.Model.Url, bytes.NewBuffer(body))
	if err != nil {
		return "", nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if cfg.Model.ApiKey != "" {
		req.Header.Add("Authorization", "Bearer "+cfg.Model.ApiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error sending request:", string(body))
	}

	if cfg.Model.ModelType == "ollama" {
		// var response ResponseBody
		var response OllamaResponse

		err = json.Unmarshal(body, &response)
		if err != nil {
			panic(err)
		}

		return response.Message.Content, nil
	}

	var response ResponseBody

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return response.Choices[0].Message.Content, nil
}

func Commit(messageContent string) error {
	commitFileName := "git_commit_editmsg"
	// Create temp file
	tmpFile, err := os.CreateTemp("", commitFileName)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(messageContent); err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}

	cmd := exec.Command("git", "commit", "-F", tmpFile.Name(), "-e")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func BuildMessages(template, issue, commitType string) ([]Message, error) {
	fileContent, err := os.ReadFile(fmt.Sprintf("prompts/%s.md", template))
	if err != nil {
		return nil, err
	}

	systemMessage := Message{
		Role:    System,
		Content: string(fileContent),
	}

	diff, err := GetGitDiff()
	if err != nil {
		return nil, err
	}

	userMessage := Message{
		Role:    User,
		Content: fmt.Sprintf("Please create a commit message based on the git diff provided below, the following fields are given:\n- **Ticket:** %s\n- **Type:** %s\n\n ## **Git Diff:**\n\n%s", issue, commitType, diff),
	}

	return []Message{
		systemMessage,
		userMessage,
	}, nil
}

func GetGitDiff() (string, error) {
	// Execute the git diff command
	cmd := exec.Command("git", "diff", "HEAD")

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Print the output
	return string(output), nil
}

func HasStagedChanges() (bool, error) {
	cmd := exec.Command("git", "diff", "--cached", "--quiet")
	err := cmd.Run()

	if err != nil {
		// If the command fails, it means there are staged changes
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode() != 0, nil
		}
		return false, err
	}

	// If the command succeeds, it means there are no staged changes
	return false, nil
}

type RequestBody struct {
	Model    Model     `json:"model,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    Role   `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type ResponseBody struct {
	Id      string
	Object  string
	Model   string
	Choices []Choice
}

type Reason string

const (
	Stop        Reason = "stop"
	Length      Reason = "length"
	ModelLength Reason = "model_length"
	Error       Reason = "error"
	ToolCalls   Reason = "tool_calls"
)

type Choice struct {
	Index        int
	Message      Message
	FinishReason Reason
}

type OllamaResponse struct {
	Message Message
}
