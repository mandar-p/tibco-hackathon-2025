package main

import (
	"bufio"
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

	log.Println("Starting command executor...")
	log.Printf("Watching for changes in: %s", cfg.FilePaths.CommandsFile)
	log.Printf("BWDesign executable: %s", cfg.BWDesign.ExecutablePath)
	log.Printf("Workspace path: %s", cfg.BWDesign.WorkspacePath)

	// Check for file changes in a loop
	var lastModTime time.Time

	for {
		// Check file modification time
		fileInfo, err := os.Stat(cfg.FilePaths.CommandsFile)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Printf("Error checking file: %v", err)
			}
			time.Sleep(5 * time.Second)
			continue
		}

		// If file was modified since last check
		if fileInfo.ModTime().After(lastModTime) {
			lastModTime = fileInfo.ModTime()
			log.Println("Commands file changed, executing commands...")
			
			if err := executeCommands(); err != nil {
				log.Printf("Error executing commands: %v", err)
			}
		}

		// Check every 5 seconds
		time.Sleep(5 * time.Second)
	}
}

func executeCommands() error {
	// Read commands from file
	file, err := os.Open(cfg.FilePaths.CommandsFile)
	if err != nil {
		return fmt.Errorf("failed to open commands file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	commands := []string{}
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			commands = append(commands, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading commands file: %v", err)
	}

	if len(commands) == 0 {
		log.Println("No commands to execute")
		return nil
	}

	// Log execution start
	logMessage := fmt.Sprintf("Starting execution of %d commands at %s\n", len(commands), time.Now().Format(time.RFC3339))
	appendToLog(logMessage)

	// Check if first command is bwdesign - if so, handle interactive session
	if len(commands) > 0 && strings.HasPrefix(commands[0], "bwdesign") {
		log.Println("Detected bwdesign command - starting interactive session")
		return executeInteractiveSession(commands)
	}

	// Fallback to individual command execution
	log.Println("No bwdesign command detected - using individual execution")
	return executeIndividualCommands(commands)
}

func executeInteractiveSession(commands []string) error {
	log.Printf("Starting BWDesign interactive session with %d commands", len(commands))

	// Parse the bwdesign command arguments
	bwdesignArgs := strings.Fields(commands[0])
	if len(bwdesignArgs) < 1 {
		return fmt.Errorf("invalid bwdesign command")
	}

	// Start BWDesign process with arguments (skip "bwdesign" itself)
	var cmdArgs []string
	if len(bwdesignArgs) > 1 {
		cmdArgs = bwdesignArgs[1:]
	}
	
	bwdesignCmd := exec.Command(cfg.BWDesign.ExecutablePath, cmdArgs...)
	bwdesignCmd.Dir = filepath.Dir(cfg.BWDesign.ExecutablePath)

	// Create pipes for stdin and stdout
	stdin, err := bwdesignCmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	stdout, err := bwdesignCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	stderr, err := bwdesignCmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	// Start the BWDesign process
	if err := bwdesignCmd.Start(); err != nil {
		return fmt.Errorf("failed to start BWDesign: %v", err)
	}

	log.Printf("BWDesign process started with PID: %d", bwdesignCmd.Process.Pid)

	// Create channels to capture output
	outputDone := make(chan bool, 2)
	
	// Read from stdout
	go func() {
		defer func() { outputDone <- true }()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("BWDesign stdout: %s", line)
			appendToLog(fmt.Sprintf("BWDesign stdout: %s\n", line))
		}
	}()

	// Read from stderr
	go func() {
		defer func() { outputDone <- true }()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("BWDesign stderr: %s", line)
			appendToLog(fmt.Sprintf("BWDesign stderr: %s\n", line))
		}
	}()

	// Wait for BWDesign to start up and show the prompt
	log.Println("Waiting for BWDesign to initialize...")
	time.Sleep(5 * time.Second)

	// Execute remaining commands in the BWDesign shell
	for i, command := range commands[1:] {
		commandIndex := i + 2
		log.Printf("Sending command %d/%d to BWDesign shell: %s", commandIndex, len(commands), command)
		
		// Send command to BWDesign stdin
		if _, err := fmt.Fprintf(stdin, "%s\n", command); err != nil {
			log.Printf("Failed to send command: %v", err)
			break
		}

		// Wait between commands to allow processing
		log.Printf("Waiting for command %d to process...", commandIndex)
		time.Sleep(3 * time.Second)
	}

	// Send exit command to close BWDesign gracefully
	log.Println("Sending exit command to BWDesign")
	fmt.Fprintf(stdin, "exit\n")
	stdin.Close()

	// Wait for the process to complete
	if err := bwdesignCmd.Wait(); err != nil {
		log.Printf("BWDesign process finished with error: %v", err)
		// Don't return error here as some BWDesign operations might exit with non-zero status
	} else {
		log.Println("BWDesign process completed successfully")
	}

	// Wait for output capture to complete (with timeout)
	outputCount := 0
	for outputCount < 2 {
		select {
		case <-outputDone:
			outputCount++
		case <-time.After(5 * time.Second):
			log.Println("Timeout waiting for output capture to complete")
			outputCount = 2 // Force exit
		}
	}

	completionMsg := fmt.Sprintf("Interactive BWDesign session completed at %s\n", time.Now().Format(time.RFC3339))
	log.Print(completionMsg)
	appendToLog(completionMsg)

	return nil
}

func executeIndividualCommands(commands []string) error {
	// This shouldn't be called when bwdesign is present, but keeping as fallback
	for i, command := range commands {
		log.Printf("Executing individual command %d/%d: %s", i+1, len(commands), command)
		
		if err := executeCommand(command); err != nil {
			errorMsg := fmt.Sprintf("Command %d failed: %s - Error: %v\n", i+1, command, err)
			log.Print(errorMsg)
			appendToLog(errorMsg)
			return err
		}

		successMsg := fmt.Sprintf("Command %d completed successfully: %s\n", i+1, command)
		log.Print(successMsg)
		appendToLog(successMsg)

		time.Sleep(2 * time.Second)
	}

	return nil
}

func executeCommand(command string) error {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return fmt.Errorf("empty command")
	}

	cmd := exec.Command(cfg.BWDesign.ExecutablePath, parts...)
	cmd.Dir = filepath.Dir(cfg.BWDesign.ExecutablePath)
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %v, output: %s", err, string(output))
	}

	if len(output) > 0 {
		outputMsg := fmt.Sprintf("Command output: %s\n", string(output))
		log.Print(outputMsg)
		appendToLog(outputMsg)
	}

	return nil
}

func appendToLog(message string) {
	file, err := os.OpenFile(cfg.FilePaths.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to write to log file: %v", err)
		return
	}
	defer file.Close()

	file.WriteString(message)
}
