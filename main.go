package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const (
	URL        = "https://api.mistral.ai/v1/chat/completions"
	KEY        = "SooCPfZithUlFcVgiKcBjk60zlKvF9nT"
	TICKET_KEY = "TICKET"
)

type Model string

const (
	MistralLargeLatest Model = "mistral-large-latest"
)

type Role string

const (
	System Role = "system"
	User   Role = "user"
)

func main() {
	msg, err := SendRequest()
	if err != nil {
		panic(err)
	}

	err = Commit(msg)
	if err != nil {
		panic(err)
	}
}

func SendRequest() (string, error) {
	ticket := os.Getenv(TICKET_KEY)
	messages, err := BuildMessages("concise", ticket, "feat")
	if err != nil {
		return "", nil
	}
	requestBody := RequestBody{
		Model:    MistralLargeLatest,
		Messages: messages,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", nil
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body))
	if err != nil {
		return "", nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+KEY)

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

func BuildMessages(template, ticket, commitType string) ([]Message, error) {
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
		Content: fmt.Sprintf("Please create a commit message based on the git diff provided below, the following fields are given:\n- **Ticket:** %s\n- **Type:** %s\n\n ## **Git Diff:**\n\n%s", ticket, commitType, diff),
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

type RequestBody struct {
	Model    Model     `json:"model,omitempty"`
	Messages []Message `json:"messages,omitempty"`
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
