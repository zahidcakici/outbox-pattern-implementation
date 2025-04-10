package handlers

import (
	"encoding/json"
	"log"
	"order-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateOrder(c *fiber.Ctx, db *gorm.DB) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	tx := db.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	outbox := models.Outbox{
		EventType: "OrderCreated",
		Payload:   toJSON(order),
		Processed: false,
	}

	if err := tx.Create(&outbox).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	tx.Commit()
	log.Printf("Order created: %v", order)
	log.Printf("Outbox entry created: %v", outbox)
	return c.Status(fiber.StatusCreated).JSON(order)
}

func toJSON(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
