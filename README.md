# 🚀 Unified LLM API Gateway

![Repository Size](https://img.shields.io/github/repo-size/JawherKl/llm-api-gateway)
![Last Commit](https://img.shields.io/github/last-commit/JawherKl/llm-api-gateway)
![Issues](https://img.shields.io/github/issues-raw/JawherKl/llm-api-gateway)
![Forks](https://img.shields.io/github/forks/JawherKl/llm-api-gateway)
![Stars](https://img.shields.io/github/stars/JawherKl/llm-api-gateway)

![Gateway Banner](https://raw.githubusercontent.com/JawherKl/llm-api-gateway/refs/heads/main/llm-api-gateway.jpg)

---

## ✨ Overview

**Unified LLM API Gateway** is a scalable, extensible platform that aggregates and normalises calls to multiple LLM backends (OpenAI, Hugging Face, Groq, Anthropic, Gemini, and more).  
It provides a unified API with built-in caching, rate limiting, authentication, logging, metrics, and production-ready deployment manifests for Docker and Kubernetes.

---

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-Framework-green.svg)](https://gin-gonic.com/)
[![OpenRouter AI](https://img.shields.io/badge/OpenAI-black.svg)](https://openai.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

## 🏗️ Architecture

- **API Gateway (Go):**  
  - Accepts client requests
  - Handles authentication, routing and request transformation
  - Aggregates/fans-out to LLM backends
- **LLM Adapters (microservices):**  
  - Wrap each provider’s API (OpenAI, Hugging Face, etc.) with a unified internal interface
- **Cache Layer:**  
  - Redis for result caching (prompt+params as cache key)
- **Rate Limiter:**  
  - Redis-based leaky-bucket or token-bucket (shared across instances)
- **Auth & Quotas:**  
  - API keys / JWT, per-key quotas (Redis or DB)
- **Observability:**  
  - Structured logs (JSON), Prometheus metrics, traces
- **Deployment:**  
  - Docker images, Helm charts, Kubernetes manifests, CI builds

---

## 📁 Monorepo Layout

```
llm-api-gateway/
├── README.md
├── LICENSE
├── .github/           # CI/CD workflows
├── infra/             # Docker Compose & Kubernetes manifests
│   ├── k8s/
│   └── docker-compose.yml
├── gateway/           # Go API gateway
│   ├── cmd/server/
│   ├── internal/
│   │   ├── handlers/
│   │   ├── adapters/
│   │   ├── cache/
│   │   ├── ratelimit/
│   │   └── metrics/
│   ├── go.mod
│   └── Dockerfile
├── adapters/          # Per-provider adapters (microservices)
│   ├── openai-adapter/
│   └── hf-adapter/
├── admin/             # NestJS admin dashboard (API keys, usage, logs)
│   ├── src/
│   ├── package.json
│   └── Dockerfile
└── tooling/
    └── tests/         # e2e test helpers
```

---

## ⚡ Quickstart

1. **Start all services:**
   ```sh
   docker-compose up --build
   ```
2. **Gateway API:**  
   [http://localhost:3020/gateway/query](http://localhost:3020/gateway/query)
3. **Admin Dashboard:**  
   [http://localhost:3040](http://localhost:3040)

---

## 🔌 Supported Providers

- [x] OpenAI (GPT-3.5, GPT-4, GPT-4o, etc.)
- [x] Hugging Face Inference API
- [x] Groq
- [x] OpenRouter
- [x] Anthropic (Claude)
- [x] Gemini (Google)
- [ ] More coming soon!

---

## 🛡️ Features

- **Unified API:** One endpoint for all LLMs
- **Authentication:** API key/JWT middleware
- **Caching:** Redis-based, prompt+params as key
- **Rate Limiting:** Per-key, Redis-backed
- **Logging:** Structured, JSON logs
- **Monitoring:** Prometheus metrics endpoint
- **Adapters:** Microservices for each provider
- **Kubernetes & Docker:** Production-ready manifests

---

## 🧑‍💻 API Usage

### Request

```http
POST /gateway/query
Authorization: <your-gateway-api-key>
Content-Type: application/json

{
  "provider": "openai" | "hf" | "groq" | "openrouter" | "anthropic" | "gemini",
  "prompt": "Your prompt here"
}
```

### Response

```json
{
  "cached": false,
  "response": "LLM output"
}
```

---

## 🚦 Development Phases

### Phase 1: Core Gateway
- [x] Unified `/query` endpoint
- [x] OpenAI, Hugging Face, Groq, OpenRouter, Anthropic, Gemini support
- [x] Redis caching
- [x] API key authentication
- [x] Rate limiting
- [x] Logging

### Phase 2: Adapters & Extensibility
- [x] Per-provider adapters as microservices
- [x] Unified internal API for adapters
- [x] Docker Compose & K8s manifests

### Phase 3: Observability & Admin
- [x] Prometheus metrics
- [x] Admin dashboard (NestJS)
- [ ] Usage quotas & billing
- [ ] Tracing (OpenTelemetry)

### Phase 4: Advanced Features (Planned)
- [ ] Multi-provider aggregation/fan-out
- [ ] Request/response transforms
- [ ] Fine-grained quotas & billing
- [ ] User/project management
- [ ] Webhooks & streaming
- [ ] Model selection & fallback
- [ ] More adapters (Cohere, Mistral, etc.)

---

## 📈 Roadmap

- [ ] Add more LLM providers & adapters
- [ ] Streaming & webhooks support
- [ ] Advanced admin features (usage, billing, analytics)
- [ ] Helm charts for K8s
- [ ] OpenAPI/Swagger docs

---

## 🤝 Contributing

Contributions are welcome! Please open issues or PRs for bugs, features, or improvements.
<!--

## 📄 License

MIT
-->
