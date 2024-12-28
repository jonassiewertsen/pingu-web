package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
)

func createSessionStore(config *Config) *session.Store {
	sessionConfig := session.Config{
		Expiration:     config.SessionExpiration,
		CookieSecure:   config.CookieSecure,
		CookieHTTPOnly: config.CookieHTTPOnly,
		CookieSameSite: config.CookieSameSite,
		KeyLookup:      config.KeyLookup,
		Storage:        createSessionStorage(config),
	}

	return session.New(sessionConfig)
}

/**
 * Create a session storage based on the configuration
 *
 * Right now, the only supported storage is memory
 */
func createSessionStorage(config *Config) fiber.Storage {
	switch config.SessionStorage {
	case "memory":
		return memory.New()
	default:
		panic("Invalid session storage")
	}
}
