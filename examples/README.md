# TIBCO Co-Pilot: BWCE Edition - Examples

This directory contains example prompts and use cases for generating TIBCO BWCE projects.

## 🎯 Integration Pattern Examples

### API Gateway Pattern
```json
{
  "user_prompt": "Create a REST API gateway that receives customer orders, validates the data, routes to inventory service, and returns confirmation with order ID"
}
```

### Data Transformation Pattern  
```json
{
  "user_prompt": "Build an ETL process that reads CSV files from FTP, transforms customer data to XML format, validates against schema, and loads into database"
}
```

### Message Routing Pattern
```json
{
  "user_prompt": "Design a message router that consumes JMS messages, filters by priority and customer type, and routes to different processing queues"
}
```

### Event Processing Pattern
```json
{
  "user_prompt": "Create a Kafka consumer that processes payment events, enriches with customer data from REST API, and publishes to notification service"
}
```

## 🏭 Enterprise Integration Examples

### Legacy System Bridge
```json
{
  "user_prompt": "Build a bridge service that exposes SOAP web services as REST APIs, handles authentication, and provides response transformation"
}
```

### File Processing Workflow
```json
{
  "user_prompt": "Create a file monitor that watches directory for incoming XML files, validates structure, splits large files, and processes in parallel"
}
```

### Database Synchronization
```json
{
  "user_prompt": "Design a sync service that reads changes from source database, applies business rules, and updates target system via REST calls"
}
```

## 🚀 Cloud-Native Microservice Examples

### Health Check Service
```json
{
  "user_prompt": "Build a health monitoring service that checks database connections, external APIs, and system resources, then exposes metrics endpoint"
}
```

### Authentication Gateway
```json
{
  "user_prompt": "Create an OAuth2 gateway that validates JWT tokens, enriches requests with user context, and forwards to downstream services"
}
```

### Event Aggregator
```json
{
  "user_prompt": "Design an event aggregator that collects metrics from multiple sources, applies windowing functions, and triggers alerts"
}
```

## 🔧 Configuration Examples

### Basic Configuration
```json
{
  "server": {
    "port": 8080,
    "host": "localhost"
  },
  "anthropic": {
    "api_key": "sk-ant-...",
    "model": "claude-sonnet-4-20250514"
  }
}
```

### Advanced Configuration
```json
{
  "bwdesign": {
    "executable_path": "/opt/tibco/bw/6.12/bin/bwdesign",
    "workspace_path": "/workspace/bwce-projects",
    "project_output_path": "/output/generated-projects"
  },
  "git": {
    "url": "https://github.com/company/bwce-projects.git",
    "token": "ghp_...",
    "branch": "generated-projects"
  }
}
```

## 📝 Command Examples

### Generate Simple REST Service
```bash
curl -X POST http://localhost:8080/api/v1/generate-commands \
  -H "Content-Type: application/json" \
  -d '{
    "user_prompt": "Create a simple REST service that accepts POST requests with customer data and returns success response"
  }'
```

### Generate Complex Integration Flow
```bash
curl -X POST http://localhost:8080/api/v1/generate-commands \
  -H "Content-Type: application/json" \
  -d '{
    "user_prompt": "Build a complex integration that reads from Kafka, enriches data with database lookup, applies business rules, transforms to different format, and publishes to multiple destinations"
  }'
```

## 🎨 Custom Patterns

### Template for Custom Integration
```json
{
  "pattern_name": "Custom Integration Template",
  "user_prompt": "Create [integration_type] that [source_action] from [source_system], [transformation_logic], and [destination_action] to [target_system]",
  "variables": {
    "integration_type": "service/process/gateway",
    "source_action": "reads/consumes/receives",
    "source_system": "database/queue/file/api",
    "transformation_logic": "validates/transforms/enriches/filters",
    "destination_action": "writes/publishes/sends",
    "target_system": "database/queue/file/api"
  }
}
```

## 🚀 Getting Started

1. **Choose a pattern** that matches your integration need
2. **Customize the prompt** with your specific requirements  
3. **Send the request** to the API server
4. **Watch the magic** as your BWCE project is generated!

## 💡 Tips for Better Results

- **Be specific** about data formats (JSON, XML, CSV)
- **Include error handling** requirements in your prompt
- **Specify protocols** (REST, SOAP, JMS, FTP, etc.)
- **Mention security** requirements (authentication, encryption)
- **Include performance** considerations (batch size, parallel processing)

---

**Ready to transform your integration ideas into TIBCO BWCE projects? Start with these examples! 🚀**
