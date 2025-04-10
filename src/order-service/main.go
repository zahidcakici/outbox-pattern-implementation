package main

import (
	"log"
	"order-service/handlers"
	"order-service/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Order{}, &models.Outbox{})

	app := fiber.New()

	app.Post("/orders", func(c *fiber.Ctx) error {
		return handlers.CreateOrder(c, db)
	})

	log.Fatal(app.Listen(":3000"))
}
