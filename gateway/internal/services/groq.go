package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type GroqRequest struct {
    Model    string              `json:"model"`
    Messages []map[string]string `json:"messages"`
}

type GroqResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
    Error struct {
        Message string `json:"message"`
    } `json:"error"`
}

func QueryGroq(apiKey, prompt string) (string, error) {
    reqBody := GroqRequest{
        Model: "openai/gpt-oss-20b",
        Messages: []map[string]string{
            {"role": "user", "content": prompt},
        },
    }
    jsonData, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
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
        var errResp GroqResponse
        json.Unmarshal(bodyBytes, &errResp)
        if errResp.Error.Message != "" {
            return "", fmt.Errorf("Groq error: %s", errResp.Error.Message)
        }
        return "", fmt.Errorf("Groq API error: %s", string(bodyBytes))
    }

    var result GroqResponse
    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return "", fmt.Errorf("failed to parse Groq response: %w", err)
    }

    if len(result.Choices) > 0 {
        return result.Choices[0].Message.Content, nil
    }
    return "", fmt.Errorf("no choices returned from Groq: %s", string(bodyBytes))
}