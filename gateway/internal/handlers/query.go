package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/JawherKl/gateway/internal/config"
    "github.com/JawherKl/gateway/internal/services"
)

type QueryRequest struct {
    Provider string `json:"provider" binding:"required"` // openai, hf
    Prompt   string `json:"prompt" binding:"required"`
}

func QueryHandler(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req QueryRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported provider"})
            return
        }

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"response": response})
    }
}
