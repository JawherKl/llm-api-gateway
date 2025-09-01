package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type GeminiPart struct {
    Text string `json:"text"`
}

type GeminiContent struct {
    Parts []GeminiPart `json:"parts"`
}

type GeminiRequest struct {
    Contents []GeminiContent `json:"contents"`
}

type GeminiResponse struct {
    Candidates []struct {
        Content struct {
            Parts []struct {
                Text string `json:"text"`
            } `json:"parts"`
        } `json:"content"`
    } `json:"candidates"`
    Error struct {
        Message string `json:"message"`
    } `json:"error"`
}

func QueryGemini(apiKey, prompt string) (string, error) {
    reqBody := GeminiRequest{
        Contents: []GeminiContent{
            {
                Parts: []GeminiPart{
                    {Text: prompt},
                },
            },
        },
    }
    jsonData, _ := json.Marshal(reqBody)

    url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-goog-api-key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        var errResp GeminiResponse
        json.Unmarshal(bodyBytes, &errResp)
        if errResp.Error.Message != "" {
            return "", fmt.Errorf("Gemini error: %s", errResp.Error.Message)
        }
        return "", fmt.Errorf("Gemini API error: %s", string(bodyBytes))
    }

    var result GeminiResponse
    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return "", fmt.Errorf("failed to parse Gemini response: %w", err)
    }

    if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
        return result.Candidates[0].Content.Parts[0].Text, nil
    }
    return "", fmt.Errorf("no candidates returned from Gemini: %s", string(bodyBytes))
}