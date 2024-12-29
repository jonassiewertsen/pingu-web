package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory"
	"github.com/gofiber/storage/redis/v2"
	"runtime"
)

func initializeCache(config *Config) fiber.Storage {
	switch config.SessionStorage {
	case "memory":
		return memory.New()
	case "redis":
		return createRedisStore(config)
	default:
		panic("Invalid cache storage")
	}
}

func createRedisStore(config *Config) fiber.Storage {
	return redis.New(redis.Config{
		Host:      "127.0.0.1",
		Port:      6379,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
}
