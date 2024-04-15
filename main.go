// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	db     *Database
	logger *zap.Logger
)

func init() {
	var err error
	_, err = initRedis()
	if err != nil {
		log.Fatal("Failed to initialize redisClient:", err)
	}
}

func initRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to Redis")
	return client, nil
}
func initDB(connStr string) {
	var err error
	db, err = NewDatabase(connStr)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
}

func main() {
	logger := InitLogger()

	initRedis()

	dbConnStr := os.Getenv("DB_CONN_STR")
	if dbConnStr == "" {
		logger.Fatal("DB_CONN_STR environment variable is not set")
	}

	initDB(dbConnStr)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	SetupRoutes(app, db)

	// Start server
	addr := ":8080"
	logger.Info("Server started", zap.String("address", addr))
	err := app.Listen(addr)
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
