# Telemetron: Unified Telemetry for Multi-Agent Systems

## Overview

**Telemetron** is a lightweight telemetry service that provides a **unified, machine-readable snapshot of system state** for multi-agent AI systems. Instead of chasing traces across fragmented observability tools, Telemetron exposes a single JSON endpoint that consolidates agent activity, runtime workload, task queues, and LLM usage into one comprehensive view.

### The Problem

Modern multi-agent systems are inherently difficult to debug in production:
- Agent reasoning traces stop at LLM invoke() calls
- Real failures occur in containerized runtimes, Kubernetes pods, or external services  
- System behavior becomes fragmented across multiple observability contexts
- Engineers spend 70% of debugging time on manual trace correlation ("data archaeology")

### The Solution

Telemetron addresses this by providing:
- **Single unified schema** for complete system state
- **Machine-first observability** designed for AI debugging tools (Claude Code, etc.)
- **State representation over trace chasing** 
- **Minimal surface area** focused purely on state exposure



## Features

### Unified System Snapshot
Single endpoint (`/system/state`) provides complete view of:
- **Agent State**: Active tasks, authorized models, deployment mapping
- **Runtime Workload**: Pod counts, resource usage, deployment status  
- **Task Queues**: Pending work, priorities, blocking relationships
- **LLM Usage**: Rate limits, costs, provider status

### Comprehensive Testing
- **81.8% coverage** in service layer
- **67.4% coverage** in repository layer
- HTTP endpoint testing with JSON validation
- Mock repositories for isolated testing

### Clean Architecture
- Separation of concerns (handlers, services, repositories, models)
- Structured logging with Zap
- Configuration management

### Developer Experience
- Swagger/OpenAPI documentation
- Test coverage reporting
- Race condition detection
- Clean commit history

## Quick Start

### Prerequisites
- Go 1.21 or higher
- Git

### Installation & Running

```bash
# Clone the repository
git clone <repository-url>
cd telemetron

# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### Basic Usage

```bash
# Get complete system state
curl http://localhost:8080/system/state

# View API documentation
open http://localhost:8080/swagger/
```

## API Documentation

### Primary Endpoint

#### `GET /system/state`

Returns unified system state as JSON conforming to the comprehensive schema.

**Response Structure:**
```json
{
  "id": "system-instance-id",
  "agents": [
    {
      "name": "agent-name",
      "description": "Agent purpose",
      "max_parallel_invocations": 10,
      "deployment_name": "k8s-deployment",
      "models": ["gpt-4", "claude-3"],
      "activity": {
        "active_task_ids": [
          {
            "id": "task-123",
            "started_on": "2026-02-06T10:00:00Z",
            "status": "running"
          }
        ],
        "updated_at": "2026-02-06T10:05:00Z"
      }
    }
  ],
  "workload": [...],
  "queues": [...],
  "litellm": [...]
}
```

### Additional Endpoints

- `GET /` - Welcome message and navigation
- `GET /swagger/` - Interactive API documentation

## Development

### Project Structure

```
telemetron/
├── cmd/server/              # Application entry point
│   ├── main.go             # Server setup and routing
│   ├── handler_test.go     # HTTP handler tests
│   └── main_test.go        # Integration tests
├── internal/
│   ├── handlers/           # HTTP request handlers  
│   ├── models/             # Data models and schemas
│   │   ├── system_state.go
│   │   └── system_state_test.go
│   ├── repositories/       # Data access layer
│   │   ├── interfaces.go   # Repository contracts
│   │   ├── mock_*.go       # Mock implementations
│   │   └── mock_test.go    # Repository tests
│   └── services/           # Business logic layer
│       ├── system_service.go
│       └── system_service_test.go
├── pkg/
│   ├── config/             # Configuration management
│   └── logger/             # Structured logging
├── docs/                   # Generated API documentation
├── scripts/                # Development scripts
└── README.md
```

### Running Tests

```bash
# Run all tests with coverage
go test -v -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run with race detection (requires CGO)
CGO_ENABLED=1 go test -race ./...
```

### Test Coverage

| Package | Coverage |
|---------|----------|
| `internal/services` | **81.8%** |
| `internal/repositories` | **67.4%** |
| `cmd/server` | **18.5%** |

### Building

```bash
# Build for current platform
go build -o telemetron cmd/server/main.go

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o telemetron-linux cmd/server/main.go
GOOS=windows GOARCH=amd64 go build -o telemetron.exe cmd/server/main.go
```

## Design Philosophy

### Core Principles

1. **Unified system snapshot over dashboards**
   - Structured JSON state rather than visual graphs
   - Machine consumption over human-oriented dashboards

2. **State representation over trace chasing**  
   - Consolidated view instead of following trace IDs across systems
   - Present runtime reality alongside agent intent

3. **Machine-first observability**
   - Primary consumer is automated debugging agents
   - Schema enforcement and clarity prioritized

4. **Minimal surface area**
   - Focus strictly on state exposure
   - Data collection and visualization out of scope

### Schema-Driven Design

The system implements a comprehensive JSON schema that captures:

- **System Identity**: Unique instance identification
- **Agent State**: Active tasks, authorized models, deployment mapping  
- **Runtime Workload**: Kubernetes-aligned pod and resource metrics
- **Task Queues**: Pending work with priorities and dependencies
- **LLM Usage**: Rate limits, costs, provider configurations

This enables **state diffing over time** for debugging agents to reason about behavioral changes.

## Configuration

### Environment Variables

```bash
# Server configuration
SERVER_PORT=8080              # Default: 8080
LOG_LEVEL=info               # Default: info (debug, info, warn, error)

# Example
export SERVER_PORT=3000
export LOG_LEVEL=debug
```

### Configuration File

The system uses `pkg/config/config.go` for centralized configuration management with sensible defaults.

## Deployment

### Docker (Future)

```dockerfile
# Dockerfile example
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o telemetron cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/telemetron .
EXPOSE 8080
CMD ["./telemetron"]
```

### Kubernetes (Future)

```yaml
# deployment.yaml example
apiVersion: apps/v1
kind: Deployment
metadata:
  name: telemetron
spec:
  replicas: 3
  selector:
    matchLabels:
      app: telemetron
  template:
    metadata:
      labels:
        app: telemetron
    spec:
      containers:
      - name: telemetron
        image: telemetron:latest
        ports:
        - containerPort: 8080
        env:
        - name: LOG_LEVEL
          value: "info"
```

## Usage Scenarios

### 1. AI-Driven Debugging

```python
import requests

# AI debugging agent queries system state
response = requests.get("http://telemetron:8080/system/state")
system_state = response.json()

# Analyze for anomalies
for agent in system_state["agents"]:
    if len(agent["activity"]["active_task_ids"]) > agent["max_parallel_invocations"]:
        print(f"Agent {agent['name']} is overloaded!")
```

### 2. Production Monitoring

```bash
# Health check script
curl -f http://telemetron:8080/system/state > /dev/null || exit 1

# State comparison for change detection  
curl http://telemetron:8080/system/state | jq . > current_state.json
diff previous_state.json current_state.json
```

### 3. Development Debugging

```bash
# Quick system overview
curl -s http://localhost:8080/system/state | jq '{
  agents: .agents | length,
  active_tasks: [.agents[].activity.active_task_ids[]] | length,
  queue_depth: [.queues[].tasks[]] | length
}'
```


## Future Improvements

- Real Kubernetes workload integration
- Message queue connectivity (Kafka, RabbitMQ)
- LiteLLM proxy integration
- Agent activity collectors
- Historical state storage
- State diff endpoints
- Webhook notifications
- Performance metrics and alerting


---

**Telemetron**: Making multi-agent systems observable, debuggable, and explainable through unified state representation.



