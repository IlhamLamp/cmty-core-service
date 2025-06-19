package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() Config {
	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	requiredEnv := []string{
		"APP_CLIENT",
		"APP_PORT",
		"APP_NAME",
		"DB_USERNAME",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_HOST",
		"CLOUDINARY_API_KEY",
		"CLOUDINARY_API_SECRET",
		"CLOUDINARY_CLOUD_NAME",
		"JWT_ACCESS_SECRET",
		"JWT_REFRESH_SECRET",
	}

	missingEnv := checkRequiredEnv(requiredEnv)
	if len(missingEnv) > 0 {
		log.Fatalf("Missing required environment variables : %v", missingEnv)
	}

	config := Config{
		AppClient:           os.Getenv("APP_CLIENT"),
		AppPort:             os.Getenv("APP_PORT"),
		AppName:             os.Getenv("APP_NAME"),
		DbUsername:          os.Getenv("DB_USERNAME"),
		DbPassword:          os.Getenv("DB_PASSWORD"),
		DbName:              os.Getenv("DB_NAME"),
		DbHost:              os.Getenv("DB_HOST"),
		CloudinaryApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
		CloudinaryCloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		JwtAccessSecret:     os.Getenv("JWT_ACCESS_SECRET"),
		JwtRefreshSecret:    os.Getenv("JWT_REFRESH_SECRET"),
	}

	log.Println("Successfully loaded .env file")

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
