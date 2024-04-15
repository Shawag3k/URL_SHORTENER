// handler.go
package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the routes for the application using Fiber.
func SetupRoutes(app *fiber.App, db *Database) {
	app.Post("/shorten", ShortenURL(db))
	app.Get("/:shortURL", RedirectURL(db))
}

// ShortenURL handles the logic for shortening a URL.
func ShortenURL(db *Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		longURL := c.FormValue("url")
		if longURL == "" {
			return c.Status(http.StatusBadRequest).SendString("Missing URL parameter")
		}

		// Generate a random short URL
		shortURL := generateRandomString(6)

		// Save the mapping of short URL to long URL in the database
		err := db.SaveURL(shortURL, longURL)
		if err != nil {
			log.Println("Failed to store URL mapping:", err)
			return c.Status(http.StatusInternalServerError).SendString("Failed to store URL mapping")
		}

		return c.SendString("Shortened URL: " + shortURL)
	}
}

// RedirectURL handles the logic for redirecting to the original URL.
func RedirectURL(db *Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortURL := c.Params("shortURL")

		// Retrieve long URL from the database
		longURL, err := db.GetOriginalURL(shortURL)
		if err != nil {
			log.Println("Failed to retrieve URL mapping:", err)
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).SendString("Short URL not found")
			}
			return c.Status(http.StatusInternalServerError).SendString("Failed to retrieve URL mapping")
		}

		return c.Redirect(longURL)
	}
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}
