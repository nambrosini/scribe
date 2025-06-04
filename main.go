package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	// Execute the git diff command
	cmd := exec.Command("git", "diff")

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing git diff: %v", err)
	}

	// Print the output
	fmt.Printf("Git Diff Output:\n%s", output)
}

func SendRequest() {
	requestBody := RequestBody{
		Model: MistralLargeLatest,
		Messages: []Message{
			{
				Role:    User,
				Content: "Could you please write me a prompt template for a commit message?",
			},
		},
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var response ResponseBody

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	// Print the response body
	fmt.Printf("Response Message: %s\n", response.Choices[0].Message.Content)
}

func NewMessage(template, ticket, commitType string) (*Message, error) {
	fileContent, err := os.ReadFile(fmt.Sprintf("templates/%s.md", template))
	if err != nil {
		return nil, err
	}

	content := string(fileContent)

	content += fmt.Sprintf("\nTicket: %s", ticket)
	content += fmt.Sprintf("\nType: %s", commitType)

	return &Message{
		Role:    User,
		Content: content,
	}, nil
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
