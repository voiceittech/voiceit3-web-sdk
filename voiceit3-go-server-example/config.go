package main

import (
	"log"
	"os"
	"strconv"
)

// Example server configuration. All values are pulled from env vars so
// the example never ships with hardcoded demo credentials. Use a .env
// loader (e.g. github.com/joho/godotenv) at startup and source
// .env.example as the template.

func mustEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Fatalf("Required environment variable %s is not set. See .env.example.", name)
	}
	return v
}

var (
	VOICEIT_API_KEY     = mustEnv("VOICEIT_THREE_API_KEY")
	VOICEIT_API_TOKEN   = mustEnv("VOICEIT_THREE_API_TOKEN")
	VOICEIT_TEST_USER_ID = mustEnv("VOICEIT_TEST_USER_ID")
	DEMO_EMAIL          = mustEnv("DEMO_EMAIL")
	DEMO_PASSWORD       = mustEnv("DEMO_PASSWORD")
	CONTENT_LANGUAGE    = envOr("CONTENT_LANGUAGE", "no-STT")
	SESSION_EXPIRATION_TIME_HOURS = envOrInt("SESSION_EXPIRATION_TIME_HOURS", 1)
)

func envOr(name, fallback string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return fallback
}

func envOrInt(name string, fallback int) int {
	if v := os.Getenv(name); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
