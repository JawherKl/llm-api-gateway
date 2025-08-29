package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "os"
    "bytes"
)

type QueryRequest struct {
    Prompt string `json:"prompt"`
}

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
}

func main() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("OPENAI_API_KEY not set")
    }

    http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
        var req QueryRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        openaiReq := OpenAIRequest{
            Model: "gpt-4",
            Messages: []map[string]string{
                {"role": "user", "content": req.Prompt},
            },
        }
        jsonData, _ := json.Marshal(openaiReq)

        openaiHttpReq, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
        openaiHttpReq.Header.Set("Authorization", "Bearer "+apiKey)
        openaiHttpReq.Header.Set("Content-Type", "application/json")

        resp, err := http.DefaultClient.Do(openaiHttpReq)
        if err != nil {
            http.Error(w, "provider error", http.StatusBadGateway)
            return
        }
        defer resp.Body.Close()

        var openaiResp OpenAIResponse
        body, _ := io.ReadAll(resp.Body)
        json.Unmarshal(body, &openaiResp)

        var output string
        if len(openaiResp.Choices) > 0 {
            output = openaiResp.Choices[0].Message.Content
        }

        json.NewEncoder(w).Encode(map[string]string{"response": output})
    })

    log.Println("OpenAI Adapter running on :3030")
    log.Fatal(http.ListenAndServe(":3030", nil))
}