package config

import (
    "log"
    "os"

	"github.com/joho/godotenv"
)

type Config struct {
    ServerPort string
    OpenAIKey  string
    HuggingFaceKey string
    RedisAddr  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default env")
	}

    cfg := &Config{
		ServerPort:     getEnv("SERVER_PORT", "3020"),
		OpenAIKey:      getEnv("OPENAI_API_KEY", ""),
		HuggingFaceKey: getEnv("HF_API_KEY", ""),
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
	}
	if cfg.OpenAIKey == "" {
		log.Println("Warning: OPENAI_API_KEY is not set")
	}
	if cfg.HuggingFaceKey == "" {
		log.Println("Warning: HF_API_KEY is not set")
	}
    return cfg
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
