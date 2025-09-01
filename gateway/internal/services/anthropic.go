package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type AnthropicMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type AnthropicRequest struct {
    Model     string            `json:"model"`
    MaxTokens int               `json:"max_tokens"`
    Messages  []AnthropicMessage `json:"messages"`
}

type AnthropicResponse struct {
    Content string `json:"content"`
    Error   struct {
        Message string `json:"message"`
    } `json:"error"`
}

func QueryAnthropic(apiKey, prompt string) (string, error) {
    reqBody := AnthropicRequest{
        Model:     "claude-sonnet-4-20250514",
        MaxTokens: 1000,
        Messages: []AnthropicMessage{
            {Role: "user", Content: prompt},
        },
    }
    jsonData, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonData))
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("anthropic-version", "2023-06-01")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        var errResp AnthropicResponse
        json.Unmarshal(bodyBytes, &errResp)
        if errResp.Error.Message != "" {
            return "", fmt.Errorf("Anthropic error: %s", errResp.Error.Message)
        }
        return "", fmt.Errorf("Anthropic API error: %s", string(bodyBytes))
    }

    // Adjust Anthropic's response as needed
    var result map[string]interface{}
    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return "", fmt.Errorf("failed to parse Anthropic response: %w", err)
    }
    // Extract the content from the first message
    if messages, ok := result["content"].(string); ok && messages != "" {
        return messages, nil
    }
    if messages, ok := result["content"].([]interface{}); ok && len(messages) > 0 {
        if msg, ok := messages[0].(map[string]interface{}); ok {
            if content, ok := msg["text"].(string); ok {
                return content, nil
            }
        }
    }
    return "", fmt.Errorf("no content returned from Anthropic: %s", string(bodyBytes))
}