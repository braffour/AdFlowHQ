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
- **Frontend:** React + Vite
- **Database:** SQLite (POC), PostgreSQL (AWS RDS for prod)
- **Secrets:** Vault (local), AWS Secrets Manager (prod)
- **AI Agent:** Gemini API (optional in workflows)

## 📦 Repository Structure

```
adflowhq-poc/
├── activities/
│   └── ads_integration.go      # Google/Facebook/CallRail call logic
├── agent/
│   └── gemini_agent.go         # Gemini prompt, tool calls, response logic
├── workflows/
│   ├── orchestrator.go         # Main workflow: lead → AI → routing
│   └── child_tasks.go          # Separate child workflows/activities for specific domains
├── worker/
│   └── main.go                 # Setup Temporal client + Worker registration
├── client/
│   └── main.go                 # CLI to start workflows (for testing/demo)
├── config/
│   └── config.go               # Vault addresses, Gemini key paths, Temporal namespace
├── vault/                      # POC Vault docker setup
├── docker-compose.yml
├── ui/                         # React/Vite frontend
│   └── src/
├── package.json
├── tests/
│   └── workflows_test.go       # Workflow-level unit tests using Temporal test suite
├── go.mod
└── go.sum
```

### 📁 Breakdown of Key Folders
- **activities/**: External API integrations for ads platforms and CallRail.
- **agent/**: AI agent logic using Gemini API — will include prompts, tool invocations, and callbacks to workflows.
- **workflows/**:
  - `orchestrator.go`: Defines the main workflow that triggers AI analysis, decision-making, and subsequent child workflows.
  - `child_tasks.go`: Contains sub-workflows or activities dedicated to ads sync, call attribution, etc.
- **worker/main.go**: Initializes Temporal Client, sets namespace and task queue, registers workflows & activities, starts polling.
- **client/main.go**: CLI tool to start a workflow with given tenant/context (e.g. `go run client/main.go --tenant=acme`), useful for demo purposes.
- **config/**: Loads config from env/Vault, including Gemini API key, Temporal server URL, secrets path.
- **vault/**: Docker Compose to run Vault locally; stores Gemini credentials and ads platform secrets.
- **ui/**: React app scaffolded via Vite. Will query Temporal workflows via CLI or HTTP shim for signals and status polling.
- **tests/**: Unit tests for workflows using `github.com/temporalio/sdk-go/testsuite`, mocking activities and the AI agent responses.

## 🔧 Getting Started (POC)

1. Clone the repo and start services:
   ```bash
   docker-compose up -d
   go run backend/main.go
   go run worker/main.go
   ```


⸻

🔧 How to run the POC
	1.	Start Temporal:
    ```bash
    temporal server start-dev --db-filename temporal.db --ui-port 8080
    ```
  2. Set Gemini key:
    ```bash
    export GEMINI_API_KEY=your-gemini-api-key
    ```
  3.	Run worker:
    ```bash
    go run worker/main.go
    ```

  4. Run the client:
    ```bash
    go run client/main.go --tenant=acme --lead=lead42
    ```
  5.	Monitor: Use Temporal Web UI at [http://localhost:8080](http://localhost:8080) to trace execution.




📌 Roadmap

 - AI-driven workflow editor (LLM-based)
 - Usage metering and alerts per tenant
 - Optional integration with Temporal Cloud
 - Full AWS deployment scripts (Terraform)


 






