package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type OpenAIRequest struct {
    Model    string              `json:"model"`
    Messages []map[string]string `json:"messages"`
}

type OpenAIResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
    Error struct {
        Message string `json:"message"`
    } `json:"error"`
}

func QueryOpenAI(apiKey, prompt string) (string, error) {
    reqBody := OpenAIRequest{
        Model: "gpt-4.1",
        Messages: []map[string]string{
            {"role": "user", "content": prompt},
        },
    }
    jsonData, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
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
        // Try to parse error message from OpenAI
        var errResp OpenAIResponse
        json.Unmarshal(bodyBytes, &errResp)
        if errResp.Error.Message != "" {
            return "", fmt.Errorf("OpenAI error: %s", errResp.Error.Message)
        }
        return "", fmt.Errorf("OpenAI API error: %s", string(bodyBytes))
    }

    var result OpenAIResponse
    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return "", fmt.Errorf("failed to parse OpenAI response: %w", err)
    }

    if len(result.Choices) > 0 {
        return result.Choices[0].Message.Content, nil
    }
    return "", fmt.Errorf("no choices returned from OpenAI: %s", string(bodyBytes))
}