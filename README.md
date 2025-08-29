# Unified LLM API Gateway

## Project summary

**Unified LLM API Gateway** a scalable API gateway that aggregates and normalizes calls to multiple LLM backends (OpenAI, HF Inference, self-hosted models), with caching, rate limiting, logging, metrics, and deployment manifests for Docker/Kubernetes.

---

# 1) High-level architecture

* **API Gateway (edge):** accepts client requests, auth, routing, request transforms, aggregator/fan-out to LLM backends.
* **LLM Adapters (microservices):** small services that wrap each provider’s API (OpenAI, Hugging Face, etc.) to expose a unified internal interface.
* **Cache layer:** Redis (result caching, cache keys based on prompt+params).
* **Rate limiter:** Redis-based leaky-bucket or token-bucket (shared across instances).
* **Auth & Quotas:** API keys / JWT + per-key quotas stored in Redis or DB.
* **Observability:** Structured logs (JSON), metrics exported in Prometheus format, traces.
* **Deployment:** Docker images, Helm charts or Kubernetes manifests, CI builds and image publishing.

---

# 2) Repo & workspace layout (monorepo)

```
llm-api-gateway/
├── README.md
├── LICENSE
├── .github/
│   └── workflows/ci.yml
├── infra/
│   ├── k8s/                  # k8s manifests or Helm charts
│   └── docker-compose.yml
├── gateway/                  # Go API gateway
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── handlers/
│   │   ├── adapters/
│   │   ├── cache/
│   │   ├── ratelimit/
│   │   └── metrics/
│   ├── go.mod
│   └── Dockerfile
├── adapters/                 # per-provider adapters
│   ├── openai-adapter/
│   └── hf-adapter/
├── admin/                    # NestJS service for keys, dashboard, logs
│   ├── src/
│   ├── package.json
│   └── Dockerfile
└── tooling/
    └── tests/                # e2e test helpers
```

# 3) Quickstart (dev)
1. `docker-compose up --build`
2. Gateway: `http://localhost:3020/v1/llm/chat`
3. Admin: `http://localhost:3040`

# 4) Components
- gateway (Go): edge gateway
- adapters: provider adapters
- admin (NestJS): API key & usage dashboard
- infra: docker-compose / k8s manifests


