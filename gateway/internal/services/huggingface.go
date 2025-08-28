package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type HuggingFaceRequest struct {
    Inputs string `json:"inputs"`
}

func QueryHuggingFace(apiKey, prompt string) (string, error) {
    reqBody := HuggingFaceRequest{Inputs: prompt}
    jsonData, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", "https://api-inference.huggingface.co/models/facebook/bart-large-cnn", bytes.NewBuffer(jsonData))
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
        return "", fmt.Errorf("HuggingFace API error: %s", string(bodyBytes))
    }

    var result []map[string]interface{}
    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return "", fmt.Errorf("failed to parse HuggingFace response: %w", err)
    }
    if len(result) > 0 {
        if generated, ok := result[0]["summary_text"].(string); ok {
            return generated, nil
        }
    }
    return "", fmt.Errorf("no summary_text in HuggingFace response: %s", string(bodyBytes))
}