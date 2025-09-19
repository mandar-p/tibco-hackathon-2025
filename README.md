# TIBCO Co-Pilot: BWCE Edition

Instantly translate your ideas into ready-to-code TIBCO BWCE projects. This AI-powered scaffolder takes your natural language prompt describing an integration process and uses Large Language Models (LLMs) to automatically generate and scaffold a complete application structure. Currently powered by Claude AI with extensible architecture for future model integrations. Eliminate manual setup and boilerplate, allowing you to go from concept to code in seconds.

## 🚀 Overview

TIBCO Co-Pilot: BWCE Edition consists of three intelligent modules that work together to transform your integration ideas into deployable TIBCO BusinessWorks Container Edition (BWCE) projects:

1. **AI Command Generator** - Natural language to TIBCO BWDesign commands via LLM integration
2. **Interactive Project Builder** - Automated BWDesign execution and project scaffolding
3. **Smart Repository Manager** - Automatic Git integration and project organization

## ✨ Key Features

- **Natural Language Processing**: Describe your integration in plain English
- **Multi-Model AI Support**: Currently powered by Claude AI, extensible for GPT, Gemini, and other LLMs
- **Intelligent Code Generation**: AI translates requirements to TIBCO commands
- **Automated Project Scaffolding**: Complete BWCE project structure generated instantly
- **Zero Manual Setup**: From idea to working project without boilerplate
- **Git Integration**: Automatic version control and project organization
- **Template-Driven**: Secure, configurable, and extensible architecture
- **Future-Ready**: Modular design supports additional AI models and integrations

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
└── examples/             # Integration patterns and examples
```

## 🛠 Prerequisites

- **Go 1.21+** - Core runtime environment
- **TIBCO BWDesign** - BusinessWorks development tools
- **Git** - Version control system
- **AI API Access** - Claude AI (current), extensible for other LLM providers
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
- **Technology**: Go HTTP server with pluggable LLM integration
- **Current Provider**: Claude AI (Anthropic) with extensible architecture
- **Purpose**: Translates natural language to BWDesign commands
- **Input**: Plain English integration requirements
- **Output**: Executable TIBCO BWDesign command sequences
- **Future Support**: OpenAI GPT, Google Gemini, local models

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

### Core Development Team

- **[Mandar Paithankar](https://github.com/mandar-p)** - Co-Creator & Lead Architect
  - *System architecture and AI integration design*
  - *Module development and deployment orchestration*
  
- **[Govardhan Saikumar](https://github.com/Govardhansaikumar)** - Co-Creator & Core Developer
  - *Project scaffolding and BWDesign integration*
  - *Repository management and automation systems*

*Both contributors have made equal and essential contributions to bringing this vision to life.*

## 🔮 AI Model Support & Extensibility

### **Current Integration**
- **Claude AI (Anthropic)**: Primary LLM for command generation
- **Optimized Models**: `claude-sonnet-4-20250514` for best performance

### **Future Roadmap**
- **OpenAI GPT**: GPT-4, GPT-4 Turbo integration
- **Google Gemini**: Gemini Pro and Ultra support  
- **Azure OpenAI**: Enterprise AI service integration
- **Local Models**: Offline model support (Ollama, LocalAI)
- **Multi-Model**: Ensemble approaches for enhanced accuracy

### **Extensible Design**
The architecture supports pluggable AI providers through:
- **Abstract interfaces** for model communication
- **Configuration-driven** provider selection
- **Unified prompt templates** across different models
- **Response normalization** for consistent output

## 📊 Configuration Reference

### **AI Provider Settings**
- `api_key`: Your AI service API key  
- `file_id`: Uploaded TIBCO documentation file reference
- `model`: Recommended model for optimal performance
- `api_url`: Configurable endpoint for different providers

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
- **Multi-Provider Security**: Consistent security across different AI services

## 📈 Performance & Scalability

- **Concurrent Processing**: Multi-threaded command execution
- **Resource Management**: Intelligent BWDesign session handling  
- **Caching Strategy**: Optimized for repeated pattern generation
- **Error Recovery**: Robust failure handling and retry mechanisms
- **Model Optimization**: Provider-specific performance tuning

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

### **Extending AI Support**
```go
// Add new AI provider
type NewAIProvider struct {
    apiKey string
    endpoint string
}

func (p *NewAIProvider) GenerateCommands(prompt string) (string, error) {
    // Implement provider-specific logic
}
```

### **Custom Integration Patterns**
- **Template System**: Add project templates for specific patterns
- **Provider Plugins**: Extend with additional AI models or tools
- **Deployment Targets**: Add support for different cloud platforms
- **Monitoring Integration**: Connect with observability platforms

## 🐛 Troubleshooting

### **Common Issues**
1. **BWDesign Path Issues**: Ensure `bwdesign.tra` file is accessible
2. **AI API Authentication**: Verify API key and service permissions  
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
- Sample integration patterns using different AI models
- Common use case templates  
- Advanced configuration scenarios
- Performance tuning guides
- Multi-provider setup examples

## 🌟 Why TIBCO Co-Pilot: BWCE Edition?

- **🚀 Speed**: From idea to working BWCE project in seconds
- **🧠 Intelligence**: Advanced AI understands complex integration patterns
- **🔧 Flexibility**: Extensible architecture supports multiple AI providers
- **💡 Innovation**: Eliminates repetitive scaffolding and boilerplate
- **🌐 Future-Ready**: Designed for the evolving AI landscape
- **👥 Community**: Built by integration experts for integration developers

## 📄 License

This project is provided for educational and development purposes. Please ensure compliance with TIBCO licensing terms for production deployments.

---

**Transform Your Integration Ideas Into Reality - One Prompt at a Time! 🚀**

*Powered by AI. Built for Developers. Ready for the Future.*
