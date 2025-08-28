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