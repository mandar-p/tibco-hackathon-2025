# TIBCOpilot

A Go-based automation tool for TIBCO BWDesign command generation, execution, and Git integration using Claude AI.

## Overview

TIBCOpilot consists of three main modules:

1. **REST API Server** - Accepts user prompts and generates TIBCO BWDesign commands using Claude AI
2. **Command Executor** - Watches for command files and executes BWDesign commands sequentially  
3. **Git Uploader** - Automatically uploads generated projects to Git repositories

## Project Structure

```
TIBCOpilot/
├── cmd/
│   ├── api-server/     # Module 1: REST API server
│   ├── executor/       # Module 2: Command executor  
│   └── git-uploader/   # Module 3: Git uploader
├── internal/
│   └── config.go       # Configuration management
├── config/
│   └── config.json     # Main configuration file
├── data/               # Runtime data files
├── go.mod              # Go module definition
├── Makefile            # Build automation
└── README.md           # This file
```

## Prerequisites

- Go 1.21 or later
- TIBCO BWDesign installed and accessible via command line
- Git installed and configured
- Claude API key from Anthropic
- Valid Git repository with push access

## Configuration

**Important: Copy the template and add your credentials:**

```bash
# Copy the template
cp config/config.json.template config/config.json

# Edit with your actual values
nano config/config.json  # or your preferred editor
```

Update `config/config.json` with your values:

```json
{
  "server": {
    "port": 8080,
    "host": "localhost"
  },
  "anthropic": {
    "api_url": "https://api.anthropic.com/v1/messages",
    "api_key": "your-anthropic-api-key-here",
    "model": "claude-sonnet-4-20250514",
    "version": "2023-06-01",
    "beta": "files-api-2025-04-14",
    "file_id": "your-tibco-documentation-file-id",
    "max_tokens": 1024
  },
  "bwdesign": {
    "executable_path": "/path/to/bwdesign",
    "workspace_path": "/path/to/workspace", 
    "project_output_path": "/path/to/projects"
  },
  "git": {
    "url": "https://github.com/yourusername/yourrepo.git",
    "token": "your-git-token-here"
  },
  "file_paths": {
    "commands_file": "./data/commands.txt",
    "log_file": "./data/execution.log"
  }
}
```

### Configuration Fields

- **server**: HTTP server settings for the REST API
- **anthropic**: Claude AI API configuration including your API key and uploaded documentation file ID
- **bwdesign**: Paths to BWDesign executable and workspace/project directories
- **git**: Git repository URL and authentication token
- **file_paths**: Local file paths for command storage and logging

## Installation & Setup

1. Clone or download the project
2. Install dependencies:
   ```bash
   make deps
   ```
3. Create required directories:
   ```bash
   make setup
   ```
4. Update configuration in `config/config.json`
5. Build all modules:
   ```bash
   make build
   ```

## Usage

### Running Individual Modules

```bash
# Start the REST API server
make run-api

# Start the command executor (in separate terminal)
make run-executor

# Run git uploader when ready to upload
make run-git
```

### Using the REST API

Send POST requests to generate BWDesign commands:

```bash
curl -X POST http://localhost:8080/api/v1/generate-commands \
     -H "Content-Type: application/json" \
     -d '{
       "user_prompt": "Give me commands to create a new bwdesign workspace, add application module, using default process, create activity timer and log",
       "api_url": "https://api.anthropic.com/v1/messages",
       "api_key": "your-api-key",
       "model_name": "claude-sonnet-4-20250514"
     }'
```

Response:
```json
{
  "status": "success", 
  "details": "bwdesign -data C:\\myWorkspace\ncreateBWApplicationModule MyApp MyProcess 1.0.0\nsystem:createActivity MyApp MyProcess com.tibco.plugin.timer.TimerEventSource\nsystem:createActivity MyApp MyProcess com.tibco.pe.core.LogActivity"
}
```

### Workflow

1. **Generate Commands**: Send request to API server with your prompt
2. **Auto Execution**: Commands are automatically written to file and picked up by executor
3. **Sequential Processing**: Executor runs each BWDesign command in order with proper delays
4. **Git Upload**: Run git uploader to push generated project to repository

## Module Details

### Module 1: REST API Server
- Accepts HTTP POST requests with user prompts
- Calls Claude AI API with TIBCO documentation context
- Returns generated BWDesign commands
- Writes commands to file for executor pickup

### Module 2: Command Executor  
- Watches `data/commands.txt` for file changes
- Executes BWDesign commands sequentially
- Handles long-running commands without timeouts
- Logs all execution details and outputs
- Runs commands in configured project directory

### Module 3: Git Uploader
- Initializes Git repository if needed
- Adds all project files to Git
- Commits changes with timestamp
- Pushes to configured remote repository
- Supports GitHub and GitLab authentication

## Development

### Building
```bash
make build          # Build all modules
make build-api      # Build API server only
make build-executor # Build executor only
make build-git      # Build git uploader only
```

### Testing
```bash
make test-api       # Test API endpoint
```

### Cleaning
```bash
make clean          # Remove build artifacts
```

## Troubleshooting

### Common Issues

1. **API Key Issues**: Ensure your Anthropic API key is valid and has sufficient credits
2. **File ID**: Make sure the TIBCO documentation file ID exists in your Claude account
3. **BWDesign Path**: Verify the BWDesign executable path is correct and accessible
4. **Git Authentication**: Ensure Git token has push permissions to the repository
5. **File Permissions**: Check that the application has read/write access to data directory

### Logs

- Executor logs: `data/execution.log`
- API server logs: Console output
- Git uploader logs: Console output

## Security Notes

- Keep your API keys and Git tokens secure
- Do not commit sensitive configuration to version control
- Use environment variables for production deployments
- Restrict file system permissions appropriately

## License

This project is provided as-is for educational and development purposes.
