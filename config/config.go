package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const (
	thirtyDaysInSeconds = 86400 * 30
)

type Config struct {
	PublicHost              string
	Port                    string
	SendPortInCallback      bool
	DatabaseURL             string
	CookiesAuthSecret       string
	CookiesAuthAgeInSeconds int
	CookiesAuthIsSecure     bool
	CookiesAuthIsHttpOnly   bool
	GithubClientID          string
	GithubClientSecret      string
	DiscordClientID         string
	DiscordClientSecret     string
	GoogleClientID          string
	GoogleClientSecret      string
	GitlabClientID          string
	GitlabClientSecret      string
}

var Envs = initConfig()

func initConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found. Using default environment variables.")
	}

	return &Config{
		PublicHost:              getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                    getEnv("PORT", "3000"),
		SendPortInCallback:      getEnvAsBool("SEND_PORT_IN_CALLBACK", true),
		DatabaseURL:             getEnv("DATABASE_URL", "file:mnemstart.db"),
		CookiesAuthSecret:       getEnv("COOKIES_AUTH_SECRET", "youllneverguesswhatthisis"),
		CookiesAuthAgeInSeconds: getEnvAsInt("COOKIES_AUTH_AGE_IN_SECONDS", thirtyDaysInSeconds),
		CookiesAuthIsSecure:     getEnvAsBool("COOKIES_AUTH_IS_SECURE", false),
		CookiesAuthIsHttpOnly:   getEnvAsBool("COOKIES_AUTH_IS_HTTP_ONLY", true),
		GithubClientID:          getEnv("GITHUB_CLIENT_ID", ""),
		GithubClientSecret:      getEnv("GITHUB_CLIENT_SECRET", ""),
		DiscordClientID:         getEnv("DISCORD_CLIENT_ID", ""),
		DiscordClientSecret:     getEnv("DISCORD_CLIENT_SECRET", ""),
		GoogleClientID:          getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret:      getEnv("GOOGLE_CLIENT_SECRET", ""),
		GitlabClientID:          getEnv("GITLAB_CLIENT_ID", ""),
		GitlabClientSecret:      getEnv("GITLAB_CLIENT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}
