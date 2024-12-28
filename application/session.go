package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
)

func createSessionStore(config *Config) *session.Store {
	var sessionStorage fiber.Storage

	// match the session storage type
	switch config.SessionStorage {
	case "memory":
		sessionStorage = createMemoryStorage()
	default:
		panic("Unknown session storage type")
	}

	sessionConfig := session.Config{
		Expiration:     config.SessionExpiration,
		CookieSecure:   config.CookieSecure,
		CookieHTTPOnly: config.CookieHTTPOnly,
		CookieSameSite: config.CookieSameSite,
		KeyLookup:      config.KeyLookup,
		Storage:        sessionStorage,
	}

	return session.New(sessionConfig)
}

func createMemoryStorage() fiber.Storage {
	return memory.New()
}
