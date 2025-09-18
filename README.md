# TIBCO Co-Pilot: BWCE Edition

Instantly translate your ideas into ready-to-code TIBCO BWCE projects. This AI-powered scaffolder takes your natural language prompt describing an integration process and uses a Large Language Model (LLM) to automatically generate and scaffold a complete application structure. Eliminate manual setup and boilerplate, allowing you to go from concept to code in seconds.

## 🚀 Overview

TIBCO Co-Pilot: BWCE Edition consists of three intelligent modules that work together to transform your integration ideas into deployable TIBCO BusinessWorks Container Edition (BWCE) projects:

1. **AI Command Generator** - Natural language to TIBCO BWDesign commands via Claude AI
2. **Interactive Project Builder** - Automated BWDesign execution and project scaffolding
3. **Smart Repository Manager** - Automatic Git integration and project organization

## ✨ Key Features

- **Natural Language Processing**: Describe your integration in plain English
- **Intelligent Code Generation**: Claude AI translates requirements to TIBCO commands
- **Automated Project Scaffolding**: Complete BWCE project structure generated instantly
- **Zero Manual Setup**: From idea to working project without boilerplate
- **Git Integration**: Automatic version control and project organization
- **Template-Driven**: Secure, configurable, and extensible architecture

## 📁 Project Structure

```
TIBCO-Co-Pilot-BWCE/
├── cmd/
│   ├── api-server/       # AI Command Generator API
│   ├── executor/         # Interactive BWDesign Executor  
│   └── git-uploader/     # Smart Repository Manager
├── internal/             # Configuration management
├── config/               # Template-based configuration
├── BWProject/            # Generated BWCE projects (auto-created)
├── data/                 # Runtime execution data
└── docs/                 # Documentation and examples
```

## 🛠 Prerequisites

- **Go 1.21+** - Core runtime environment
- **TIBCO BWDesign** - BusinessWorks development tools
- **Git** - Version control system
- **Claude API Access** - Anthropic AI service subscription
- **GitHub/GitLab Account** - Repository hosting (optional)

## ⚡ Quick Start

### 1. **Setup Configuration**

```bash
# Copy configuration template
cp config/config.json.template config/config.json

# Edit with your credentials and paths
nano config/config.json
```

**Required Configuration:**
```json
{
  "anthropic": {
    "api_key": "your-claude-api-key",
    "file_id": "your-tibco-documentation-file-id"
  },
  "bwdesign": {
    "executable_path": "/path/to/bwdesign",
    "workspace_path": "/path/to/workspace",
    "project_output_path": "/path/to/output"
  },
  "git": {
    "url": "https://github.com/yourusername/yourrepo.git",
    "token": "your-git-token"
  }
}
```

### 2. **Build & Launch**

```bash
# Install dependencies and build
make deps && make build

# Start the AI Command Generator (Terminal 1)
./bin/api-server

# Start the Project Builder (Terminal 2)
./bin/executor
```

### 3. **Generate Your First Project**

```bash
# Send natural language prompt to AI
curl -X POST http://localhost:8080/api/v1/generate-commands \
     -H "Content-Type: application/json" \
     -d '{
       "user_prompt": "Create a REST API that receives customer data, validates it, transforms it to XML, and sends it to a legacy system via FTP",
       "api_url": "https://api.anthropic.com/v1/messages",
       "api_key": "your-api-key",
       "model_name": "claude-sonnet-4-20250514"
     }'
```

### 4. **Deploy to Repository**

```bash
# Upload generated project to Git
./bin/git-uploader
```

Your BWCE project will be automatically created in the `BWProject/` directory and pushed to your repository!

## 🎯 Use Cases

### **Enterprise Integration Patterns**
- **API Gateway**: "Create a REST API gateway that routes requests to multiple backend services"
- **Data Transformation**: "Build an ETL process that reads CSV files, transforms data, and loads into database"
- **Message Routing**: "Design a message router that processes orders and sends them to different systems based on priority"

### **Cloud-Native Microservices**
- **Event Processing**: "Create a microservice that consumes Kafka messages and processes payment events"
- **Service Orchestration**: "Build an orchestration service that coordinates multiple API calls for user onboarding"
- **Health Monitoring**: "Design a health check service that monitors system status and sends alerts"

### **Legacy System Integration**
- **Protocol Bridging**: "Create a bridge between REST APIs and SOAP web services"
- **File Processing**: "Build a file processor that monitors directories and processes incoming data files"
- **Database Synchronization**: "Design a sync service that keeps two databases in sync"

## 🏗 Architecture

### **Module 1: AI Command Generator**
- **Technology**: Go HTTP server with Claude AI integration
- **Purpose**: Translates natural language to BWDesign commands
- **Input**: Plain English integration requirements
- **Output**: Executable TIBCO BWDesign command sequences

### **Module 2: Interactive Project Builder**
- **Technology**: Go file watcher with process orchestration
- **Purpose**: Executes BWDesign commands in interactive sessions
- **Process**: Monitors command files and builds complete BWCE projects
- **Output**: Ready-to-deploy TIBCO BusinessWorks applications

### **Module 3: Smart Repository Manager**
- **Technology**: Git automation with intelligent organization
- **Purpose**: Manages version control and project structure
- **Features**: Automatic commits, branch management, and project organization
- **Integration**: Seamless GitHub/GitLab connectivity

## 🤝 Contributors

- **[Mandar Paithankar](https://github.com/mandar-p)** - Lead Developer & Architect
- **[Govardhan Saikumar](https://github.com/Govardhansaikumar)** - Core Contributor

## 📊 Configuration Reference

### **Anthropic AI Settings**
- `api_key`: Your Claude API key from Anthropic Console
- `file_id`: Uploaded TIBCO documentation file reference
- `model`: Recommended `claude-sonnet-4-20250514` for optimal performance

### **BWDesign Integration**
- `executable_path`: Full path to BWDesign binary
- `workspace_path`: TIBCO workspace directory
- `project_output_path`: Generated project destination

### **Git Repository Management**
- `url`: Target repository URL (GitHub/GitLab/etc.)
- `token`: Personal access token with repository permissions

## 🛡 Security & Best Practices

- **Template Configuration**: Secrets never committed to version control
- **Token Management**: Secure credential handling and rotation
- **Process Isolation**: Sandboxed command execution environment
- **Audit Logging**: Complete execution trace and debugging information

## 📈 Performance & Scalability

- **Concurrent Processing**: Multi-threaded command execution
- **Resource Management**: Intelligent BWDesign session handling  
- **Caching Strategy**: Optimized for repeated pattern generation
- **Error Recovery**: Robust failure handling and retry mechanisms

## 🔧 Development & Customization

### **Build Commands**
```bash
make build          # Build all modules
make build-api      # Build API server only
make build-executor # Build executor only
make build-git      # Build git uploader only
make clean          # Remove build artifacts
make test-api       # Test API endpoint
```

### **Extending Functionality**
- **Custom Templates**: Add project templates for specific patterns
- **Integration Plugins**: Extend with additional AI models or tools
- **Deployment Targets**: Add support for different cloud platforms
- **Monitoring Integration**: Connect with observability platforms

## 🐛 Troubleshooting

### **Common Issues**
1. **BWDesign Path Issues**: Ensure `bwdesign.tra` file is accessible
2. **API Authentication**: Verify Claude API key and file permissions  
3. **Git Connectivity**: Check repository URL and token validity
4. **Process Timeouts**: Adjust timing for complex project generation

### **Debug Mode**
```bash
# Enable verbose logging
export DEBUG=true
./bin/executor
```

### **Log Analysis**
- **Execution Logs**: `data/execution.log`
- **API Server Logs**: Console output with timestamps
- **Git Operations**: Real-time command output and status

## 📚 Examples & Templates

See the `/examples` directory for:
- Sample integration patterns
- Common use case templates  
- Advanced configuration scenarios
- Performance tuning guides

## 📄 License

This project is provided for educational and development purposes. Please ensure compliance with TIBCO licensing terms for production deployments.

---

**Transform Your Integration Ideas Into Reality - One Prompt at a Time! 🚀**
