package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	requiredEnv := []string{
		"AppClient",
		"AppPort",
		"AppName",
		"CloudinaryApiKey",
		"CloudinaryApiSecret",
		"CloudinaryCloudName",
		"JwtAccessSecret",
		"JwtRefreshSecret",
	}

	missingEnv := checkRequiredEnv(requiredEnv)
	if len(missingEnv) > 0 {
		log.Fatalf("Missing required environment variables : %v", missingEnv)
	}

	config := Config{
		AppClient:           os.Getenv("AppClient"),
		AppPort:             os.Getenv("AppPort"),
		AppName:             os.Getenv("AppName"),
		CloudinaryApiKey:    os.Getenv("CloudinaryApiKey"),
		CloudinaryApiSecret: os.Getenv("CloudinaryApiSecret"),
		CloudinaryCloudName: os.Getenv("CloudinaryCloudName"),
		JwtAccessSecret:     os.Getenv("JwtAccessSecret"),
		JwtRefreshSecret:    os.Getenv("JwtRefreshSecret"),
	}

	return config
}

func checkRequiredEnv(keys []string) []string {
	var missing []string
	for _, key := range keys {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}
	return missing
}
