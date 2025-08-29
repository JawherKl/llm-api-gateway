1. Set Environment Variables
Create a .env file in gateway (or set env vars directly). Example .env:
```bash
SERVER_PORT=8080
OPENAI_API_KEY=your-openai-key
HF_API_KEY=your-huggingface-key
REDIS_ADDR=localhost:6379
GATEWAY_API_KEY=your-gateway-api-key
```

2. Run the Gateway
From the gateway directory, run:

```bash
go run ./cmd/server/main.go
```

Or, from the project root:

```bash
go run ./gateway/cmd/server/main.go
```

3. Using Docker (optional)
If you prefer Docker:

```bash
docker build -t gateway ./gateway
docker run -p 3020:3020 --env-file ./gateway/.env gateway
```

Or with Docker Compose (from project root):

```bash
docker-compose up --build
```
Or
```bash
docker-compose up -d
```

4. Test the API
Send a POST request to your endpoint (e.g., using curl):

```bash
curl -X POST http://localhost:3020/gateway/query \
-H "Authorization: your_gateway_api_key" \
-H "Content-Type: application/json" \
-d '{
  "provider": "openai",
  "prompt": "Your prompt here"
}'
