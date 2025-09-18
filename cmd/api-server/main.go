package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	config "tibcopilot/internal"
)

// Request structure for incoming API calls
type GenerateRequest struct {
	UserPrompt string `json:"user_prompt"`
	APIURL     string `json:"api_url"`
	APIKey     string `json:"api_key"`
	ModelName  string `json:"model_name"`
}

// Response structure for API responses
type GenerateResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}

// Claude API structures
type ClaudeRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	System    string    `json:"system"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type   string      `json:"type"`
	Text   string      `json:"text,omitempty"`
	Source *FileSource `json:"source,omitempty"`
}

type FileSource struct {
	Type   string `json:"type"`
	FileID string `json:"file_id"`
}

type ClaudeResponse struct {
	Content []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"content"`
}

var cfg *config.Config

func main() {
	var err error
	cfg, err = config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	http.HandleFunc("/api/v1/generate-commands", handleGenerateCommands)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleGenerateCommands(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Call Claude API
	commands, err := callClaudeAPI(req)
	if err != nil {
		log.Printf("Claude API error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write commands to file
	if err := writeCommandsToFile(commands); err != nil {
		log.Printf("File write error: %v", err)
	}

	// Send response
	response := GenerateResponse{
		Status:  "success",
		Details: commands,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func callClaudeAPI(req GenerateRequest) (string, error) {
	// Prepare Claude API request
	claudeReq := ClaudeRequest{
		Model:     req.ModelName,
		MaxTokens: cfg.Anthropic.MaxTokens,
		System:    "You are a TIBCO BWDesign command generator. You MUST output ONLY executable commands from the attached TIBCO documentation. NO explanations, comments, metadata or extra text. NO code blocks, backticks or formatting. Each command on a separate line. Use exact syntax from documentation. Never invent or modify commands. If asked to explain, provide only what the command does from documentation. For generation requests, output commands only. Format: One command per line, nothing else.",
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "document",
						Source: &FileSource{
							Type:   "file",
							FileID: cfg.Anthropic.FileID,
						},
					},
					{
						Type: "text",
						Text: req.UserPrompt,
					},
				},
			},
		},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(claudeReq)
	if err != nil {
		return "", err
	}

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", req.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// Set headers
	httpReq.Header.Set("x-api-key", req.APIKey)
	httpReq.Header.Set("anthropic-version", cfg.Anthropic.Version)
	httpReq.Header.Set("anthropic-beta", cfg.Anthropic.Beta)
	httpReq.Header.Set("Content-Type", "application/json")

	// Make request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse response
	var claudeResp ClaudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		return "", err
	}

	if len(claudeResp.Content) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	return claudeResp.Content[0].Text, nil
}

func writeCommandsToFile(commands string) error {
	return os.WriteFile(cfg.FilePaths.CommandsFile, []byte(commands), 0644)
}
