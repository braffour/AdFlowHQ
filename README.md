# AdFlowHQ

AdFlowHQ is a multi-tenant, AI-enabled workflow orchestration platform built on Temporal (Go SDK), designed for automating lead journeys across Google Ads, Facebook Ads, and CallRail. It supports smart decision-making using AI agents and offers tenant-specific control via a modern React UI.

## 🌟 Key Features

- 🧩 **Workflow Orchestration** – Powered by Temporal.io with pluggable activities
- 👥 **Multi-Tenant SaaS** – Each client (tenant) has isolated data and workflows
- 🤖 **AI Agent Collaboration** – Integrate Gemini API to enrich workflows with intent analysis and smart routing
- 🔐 **Secure Secret Management** – Uses HashiCorp Vault (POC) and AWS Secrets Manager (Prod)
- 📊 **FinOps Ready** – Track and enforce cost policies per tenant and workflow
- 📞 **Marketing Integrations**
  - Google Ads: Conversion upload, audience updates
  - Facebook Ads: Lead syncing, pixel events
  - CallRail: Call tracking and lead attribution
- 🌐 **React Frontend** – Workflow management UI per tenant

## 🚀 Tech Stack

- **Backend:** Go + Temporal SDK
- **Frontend:** React + Vite + TypeScript + Tailwind CSS
- **Database:** PostgreSQL (via Docker)
- **Secrets:** Vault (local), AWS Secrets Manager (prod)
- **AI Agent:** Gemini API (Google AI)
- **Containerization:** Docker + Docker Compose

## 📦 Repository Structure

```
AdFlowHQ/
├── activities/
│   └── ads_integration.go      # Google/Facebook/CallRail API integrations
├── agent/
│   └── gemini_agent.go         # Gemini AI agent implementation
├── workflows/
│   └── orchestrator.go         # Main workflow orchestration
├── worker/
│   └── main.go                 # Temporal worker setup
├── client/
│   └── main.go                 # CLI client for testing workflows
├── config/
│   └── config.go               # Configuration management
├── ui/                         # React frontend
│   ├── src/
│   │   ├── components/         # Reusable UI components
│   │   ├── pages/             # Page components
│   │   ├── App.tsx            # Main app component
│   │   └── main.tsx           # App entry point
│   ├── package.json           # Frontend dependencies
│   ├── vite.config.ts         # Vite configuration
│   └── tailwind.config.js     # Tailwind CSS configuration
├── vault/                      # Vault configuration
├── docker-compose.yml          # Multi-service orchestration
├── Dockerfile.worker           # Worker container
├── Dockerfile.client           # Client container
├── Makefile                    # Development commands
├── go.mod                      # Go module definition
└── README.md                   # This file
```

## 🔧 Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+
- Node.js 18+
- Gemini API key (optional, for AI features)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd AdFlowHQ

# Install Go dependencies
make deps

# Install UI dependencies
make ui-deps
```

### 2. Environment Setup

Set your Gemini API key (optional):
```bash
export GEMINI_API_KEY=your-gemini-api-key
```

### 3. Start Development Environment

```bash
# Start all services (Temporal, Vault, PostgreSQL)
make dev
```

This will:
- Start Temporal server with PostgreSQL backend
- Start Vault for secrets management
- Start the AdFlowHQ worker
- Setup initial Vault secrets
- Start the React UI

### 4. Access Services

- **Temporal UI**: http://localhost:8080
- **Vault UI**: http://localhost:8200 (Token: `dev-token`)
- **React UI**: http://localhost:3000

## 🛠 Development

### Running Individual Services

```bash
# Run worker only
make worker

# Run client to test workflows
make client

# Start UI development server
make ui-dev

# View logs
make docker-logs
```

### Testing Workflows

```bash
# Start a workflow
go run client/main.go --tenant=acme --lead=lead123

# Monitor in Temporal UI
open http://localhost:8080
```

### Building for Production

```bash
# Build Go binaries
make build

# Build UI for production
cd ui && npm run build

# Build Docker images
docker-compose build
```

## 🔐 Secrets Management

The application uses HashiCorp Vault for secrets management. In development, secrets are automatically configured via the `make vault-setup` command.

### Required Secrets

- `secret/adflowhq/google-ads/google_ads_key`
- `secret/adflowhq/facebook-ads/facebook_ads_key`
- `secret/adflowhq/callrail/api_key`

### Adding Secrets Manually

```bash
# Using Vault CLI
vault kv put secret/adflowhq/google-ads google_ads_key="your-key"

# Using HTTP API
curl -X POST -H "X-Vault-Token: dev-token" \
  -H "Content-Type: application/json" \
  -d '{"data":{"google_ads_key":"your-key"}}' \
  http://localhost:8200/v1/secret/data/adflowhq/google-ads
```

## 🧪 Testing

```bash
# Run Go tests
make test

# Run UI tests
cd ui && npm test
```

## 📊 Monitoring

- **Temporal UI**: Monitor workflow executions, history, and performance
- **Vault UI**: Manage secrets and access policies
- **React Dashboard**: View workflow statistics and recent activity

## 🚀 Deployment

### Production Considerations

1. **Database**: Use managed PostgreSQL (AWS RDS, GCP Cloud SQL)
2. **Secrets**: Use AWS Secrets Manager or Azure Key Vault
3. **Temporal**: Consider Temporal Cloud for production workloads
4. **Monitoring**: Add Prometheus/Grafana for metrics
5. **Logging**: Implement structured logging with ELK stack

### Docker Deployment

```bash
# Build and deploy
docker-compose -f docker-compose.prod.yml up -d
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

1. **Temporal connection failed**: Ensure Temporal server is running on port 7233
2. **Vault connection failed**: Check Vault is running and token is valid
3. **UI not loading**: Verify Node.js dependencies are installed
4. **Workflow failures**: Check Temporal UI for detailed error messages

### Logs

```bash
# View all service logs
make docker-logs

# View specific service logs
docker-compose logs -f worker
docker-compose logs -f ui
```

### Reset Development Environment

```bash
# Stop and remove all containers
make docker-down

# Clean build artifacts
make clean

# Restart fresh
make dev
```






