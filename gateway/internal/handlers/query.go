package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/JawherKl/gateway/internal/config"
    "github.com/JawherKl/gateway/internal/services"
    "github.com/JawherKl/gateway/internal/cache"
    "context"
    "crypto/sha256"
    "encoding/hex"
    "time"
    "log"
)

type QueryRequest struct {
    Provider string `json:"provider" binding:"required"` // openai, hf
    Prompt   string `json:"prompt" binding:"required"`
}

func QueryHandler(cfg *config.Config, cache *cache.Cache) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req QueryRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Create a cache key based on provider+prompt
        hash := sha256.Sum256([]byte(req.Provider + ":" + req.Prompt))
        cacheKey := "llm:" + hex.EncodeToString(hash[:])

        ctx := context.Background()
        if cached, err := cache.Get(ctx, cacheKey); err == nil && cached != "" {
            log.Printf("Cache hit for %s", cacheKey)
            c.JSON(http.StatusOK, gin.H{"response": cached, "cached": true})
            return
        }

        var response string
        var err error
        switch req.Provider {
        case "openai":
            response, err = services.QueryOpenAI(cfg.OpenAIKey, req.Prompt)
        case "hf":
            response, err = services.QueryHuggingFace(cfg.HuggingFaceKey, req.Prompt)
        case "groq":
            response, err = services.QueryGroq(cfg.GroqKey, req.Prompt)
        case "openrouter":
            response, err = services.QueryOpenRouter(cfg.OpenRouterKey, req.Prompt)
        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported provider"})
            return
        }

        if err != nil {
            log.Printf("Provider error: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        cache.Set(ctx, cacheKey, response, 10*time.Minute)

        log.Printf("Cache miss for %s, storing response", cacheKey)
        cache.Set(ctx, cacheKey, response, 10*time.Minute)
        c.JSON(http.StatusOK, gin.H{"response": response, "cached": false})
    }
}
