package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type OpenRouterRequest struct {
    Model    string              `json:"model"`
    Messages []map[string]string `json:"messages"`
}

type OpenRouterResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
    Error struct {
        Message string `json:"message"`
    } `json:"error"`
}

func QueryOpenRouter(apiKey, prompt string) (string, error) {
    reqBody := OpenRouterRequest{
        Model: "openai/gpt-4o-mini",
        Messages: []map[string]string{
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": prompt},
        },
    }
    jsonData, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        var errResp OpenRouterResponse
        json.Unmarshal(bodyBytes, &errResp)
        if errResp.Error.Message != "" {
            return "", fmt.Errorf("OpenRouter error: %s", errResp.Error.Message)
        }
        return "", fmt.Errorf("OpenRouter API error: %s", string(bodyBytes))
    }

    var result OpenRouterResponse
    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return "", fmt.Errorf("failed to parse OpenRouter response: %w", err)
    }

    if len(result.Choices) > 0 {
        return result.Choices[0].Message.Content, nil
    }
    return "", fmt.Errorf("no choices returned from OpenRouter: %s", string(bodyBytes))
}