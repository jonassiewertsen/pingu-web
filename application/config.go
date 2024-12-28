package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"time"
)

type Config struct {
	AppEnv   string
	AppPort  string
	AppIsDev bool

	Fiber fiber.Config

	SessionExpiration time.Duration
	SessionStorage    string
	CookieSecure      bool
	CookieHTTPOnly    bool
	CookieSameSite    string
	KeyLookup         string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	config := &Config{
		AppEnv:            os.Getenv("APP_ENV"),
		AppPort:           os.Getenv("SERVER_PORT"),
		Fiber:             fiber.Config{Views: nil},
		SessionExpiration: 24 * time.Hour,
		SessionStorage:    "memory", // Supporte: memory
		CookieSecure:      false,
		CookieHTTPOnly:    true,
		CookieSameSite:    "Lax",
		KeyLookup:         "cookie:pingu_web",
	}

	setDefaults(config)
	validate(config)

	return config
}

func setDefaults(config *Config) {
	if config.AppEnv == "" {
		config.AppEnv = "development"
	}

	if config.AppEnv == "development" {
		config.AppIsDev = true
	}

	if config.AppPort == "" {
		config.AppPort = ":8080"
	}
}

func validate(config *Config) {
	if config.AppEnv != "development" && config.AppEnv != "production" {
		panic("Invalid APP_ENV value. Must be 'development' or 'production'")
	}

	if config.AppPort == "" {
		panic("Invalid SERVER_PORT value. Must be a valid port number")
	}
	// Check if string starts with : and then a number
	if !strings.HasPrefix(config.AppPort, ":") {
		panic("Invalid SERVER_PORT value. Must be a valid port number and start with ':'. Fx. :8080")
	}
}
