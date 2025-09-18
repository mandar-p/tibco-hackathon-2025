package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	config "tibcopilot/internal"
)

var cfg *config.Config

func main() {
	var err error
	cfg, err = config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	log.Println("Starting Git uploader...")

	if err := uploadToGit(); err != nil {
		log.Fatal("Failed to upload to Git:", err)
	}

	log.Println("Successfully uploaded project to Git")
}

func uploadToGit() error {
	projectPath := cfg.BWDesign.ProjectOutputPath
	gitURL := cfg.Git.URL
	gitToken := cfg.Git.Token

	// Check if project directory exists
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return fmt.Errorf("project directory does not exist: %s", projectPath)
	}

	log.Printf("Uploading project from: %s", projectPath)
	log.Printf("Git URL: %s", gitURL)

	// Change to project directory
	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %v", err)
	}

	// Check if git is already initialized
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Println("Initializing Git repository...")
		if err := runGitCommand("init"); err != nil {
			return fmt.Errorf("failed to initialize git: %v", err)
		}
	}

	// Add all files
	log.Println("Adding files to Git...")
	if err := runGitCommand("add", "."); err != nil {
		return fmt.Errorf("failed to add files: %v", err)
	}

	// Check if there are any changes to commit
	output, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %v", err)
	}

	if len(strings.TrimSpace(string(output))) == 0 {
		log.Println("No changes to commit")
		return nil
	}

	// Commit changes
	commitMessage := fmt.Sprintf("Auto-commit from TIBCOpilot at %s", time.Now().Format("2006-01-02 15:04:05"))
	log.Printf("Committing changes: %s", commitMessage)
	if err := runGitCommand("commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("failed to commit: %v", err)
	}

	// Set remote origin if not exists
	remoteURL := buildAuthenticatedURL(gitURL, gitToken)
	if err := setRemoteOrigin(remoteURL); err != nil {
		return fmt.Errorf("failed to set remote origin: %v", err)
	}

	// Push to remote
	log.Println("Pushing to remote repository...")
	if err := runGitCommand("push", "-u", "origin", "main"); err != nil {
		// Try with master if main fails
		if err := runGitCommand("push", "-u", "origin", "master"); err != nil {
			return fmt.Errorf("failed to push to remote: %v", err)
		}
	}

	return nil
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func setRemoteOrigin(remoteURL string) error {
	// Check if origin already exists
	cmd := exec.Command("git", "remote", "get-url", "origin")
	if err := cmd.Run(); err != nil {
		// Origin doesn't exist, add it
		log.Println("Adding remote origin...")
		return runGitCommand("remote", "add", "origin", remoteURL)
	}

	// Origin exists, update it
	log.Println("Updating remote origin...")
	return runGitCommand("remote", "set-url", "origin", remoteURL)
}

func buildAuthenticatedURL(gitURL, token string) string {
	// Handle different Git URL formats
	if strings.HasPrefix(gitURL, "https://github.com/") {
		// GitHub URL - inject token
		parts := strings.SplitN(gitURL, "://", 2)
		if len(parts) == 2 {
			return fmt.Sprintf("https://%s@%s", token, parts[1])
		}
	} else if strings.HasPrefix(gitURL, "https://gitlab.com/") {
		// GitLab URL - inject token
		parts := strings.SplitN(gitURL, "://", 2)
		if len(parts) == 2 {
			return fmt.Sprintf("https://oauth2:%s@%s", token, parts[1])
		}
	}

	// For other Git providers or if parsing fails, return original URL
	// User should ensure the URL already contains authentication
	log.Printf("Using original Git URL (ensure it contains authentication): %s", gitURL)
	return gitURL
}
