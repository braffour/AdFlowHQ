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

/
├── backend/ # Go API and workflow activities
├── worker/ # Temporal worker logic
├── frontend/ # React-based multi-tenant UI
├── db/ # Schema and migration files
├── docker/ # Local environment setup
└── README.md


## 🔧 Getting Started (POC)

1. Clone the repo and start services:
   ```bash
   docker-compose up -d

go run backend/main.go
go run worker/main.go

cd frontend && npm install && npm run dev

📌 Roadmap

 AI-driven workflow editor (LLM-based)
 Usage metering and alerts per tenant
 Optional integration with Temporal Cloud
 Full AWS deployment scripts (Terraform)






