## Gateway API

### POST /gateway/query

**Request:**
```json
{
  "provider": "openai" | "hf",
  "prompt": "Your prompt here"
}
```

**Response:**
```json
{
  "response": "LLM output"
}
```

Create your `.env` file in the `gateway/` directory with your API keys and Redis address:

```
SERVER_PORT=your_port
OPENAI_API_KEY=your_openai_api_key
HF_API_KEY=your_hf_api_key
REDIS_ADDR=localhost:6379
GATEWAY_API_KEY=your_gateway_api_key
```