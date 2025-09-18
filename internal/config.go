package config

import (
	"encoding/json"
	"os"
)

// Config represents the application configuration
type Config struct {
	Server struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
	Anthropic struct {
		APIURL    string `json:"api_url"`
		APIKey    string `json:"api_key"`
		Model     string `json:"model"`
		Version   string `json:"version"`
		Beta      string `json:"beta"`
		FileID    string `json:"file_id"`
		MaxTokens int    `json:"max_tokens"`
	} `json:"anthropic"`
	BWDesign struct {
		ExecutablePath    string `json:"executable_path"`
		WorkspacePath     string `json:"workspace_path"`
		ProjectOutputPath string `json:"project_output_path"`
	} `json:"bwdesign"`
	Git struct {
		URL   string `json:"url"`
		Token string `json:"token"`
	} `json:"git"`
	FilePaths struct {
		CommandsFile string `json:"commands_file"`
		LogFile      string `json:"log_file"`
	} `json:"file_paths"`
}

// LoadConfig loads configuration from JSON file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
