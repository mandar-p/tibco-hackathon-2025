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

	log.Printf("Uploading BWCE project from: %s", projectPath)
	log.Printf("Git URL: %s", gitURL)

	// Get the current working directory (should be the TIBCOpilot repo root)
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	// Create BWProject subdirectory in the current repo
	bwProjectDir := filepath.Join(currentDir, "BWProject")
	if err := os.MkdirAll(bwProjectDir, 0755); err != nil {
		return fmt.Errorf("failed to create BWProject directory: %v", err)
	}

	log.Printf("Created BWProject directory: %s", bwProjectDir)

	// Copy all files from project directory to BWProject subdirectory
	if err := copyDirectory(projectPath, bwProjectDir); err != nil {
		return fmt.Errorf("failed to copy project files: %v", err)
	}

	log.Println("Copied BWCE project files to BWProject subdirectory")

	// Change to the repo root directory
	if err := os.Chdir(currentDir); err != nil {
		return fmt.Errorf("failed to change to repo directory: %v", err)
	}

	// Check if git is already initialized
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Println("Initializing Git repository...")
		if err := runGitCommand("init"); err != nil {
			return fmt.Errorf("failed to initialize git: %v", err)
		}
	}

	// Add BWProject files to Git
	log.Println("Adding BWProject files to Git...")
	if err := runGitCommand("add", "BWProject/"); err != nil {
		return fmt.Errorf("failed to add BWProject files: %v", err)
	}

	// Check if there are any changes to commit
	output, err := exec.Command("git", "status", "--porcelain", "BWProject/").Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %v", err)
	}

	if len(strings.TrimSpace(string(output))) == 0 {
		log.Println("No changes to commit in BWProject")
		return nil
	}

	// Commit changes
	commitMessage := fmt.Sprintf("Add generated BWCE project at %s", time.Now().Format("2006-01-02 15:04:05"))
	log.Printf("Committing BWCE project: %s", commitMessage)
	if err := runGitCommand("commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("failed to commit: %v", err)
	}

	// Set remote origin if not exists
	remoteURL := buildAuthenticatedURL(gitURL, gitToken)
	if err := setRemoteOrigin(remoteURL); err != nil {
		return fmt.Errorf("failed to set remote origin: %v", err)
	}

	// Push to remote
	log.Println("Pushing BWCE project to remote repository...")
	if err := runGitCommand("push", "origin", "main"); err != nil {
		// Try with master if main fails
		if err := runGitCommand("push", "origin", "master"); err != nil {
			return fmt.Errorf("failed to push to remote: %v", err)
		}
	}

	return nil
}

func copyDirectory(srcDir, destDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from source directory
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// Skip hidden files and directories
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// Copy file
		return copyFile(path, destPath)
	})
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create destination directory if needed
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy file contents
	buf := make([]byte, 4096)
	for {
		n, err := sourceFile.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destFile.Write(buf[:n]); err != nil {
			return err
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
