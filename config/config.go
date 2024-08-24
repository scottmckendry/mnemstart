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
}

var Envs = initConfig()

func initConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
		log.Println("Creating empty .env file, with default values for required variables")
		createEmptyEnvFile()
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
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
		GithubClientID:          getEnvOrPanic("GITHUB_CLIENT_ID"),
		GithubClientSecret:      getEnvOrPanic("GITHUB_CLIENT_SECRET"),
		DiscordClientID:         getEnvOrPanic("DISCORD_CLIENT_ID"),
		DiscordClientSecret:     getEnvOrPanic("DISCORD_CLIENT_SECRET"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvOrPanic(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("Missing required environment variable: " + key)
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

func createEmptyEnvFile() {
	f, err := os.Create(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(
		"GITHUB_CLIENT_ID=\nGITHUB_CLIENT_SECRET=\nDISCORD_CLIENT_ID=\nDISCORD_CLIENT_SECRET=\n",
	)
	if err != nil {
		log.Fatal(err)
	}
}
