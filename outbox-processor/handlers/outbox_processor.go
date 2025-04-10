package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"outbox-processor/models"
	"time"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func Start(ctx context.Context, db *gorm.DB) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-events",
	})

	defer writer.Close()

	fmt.Println("Outbox processor started")

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down outbox processor...")
			return ctx.Err()
		default:
			var outboxes []models.Outbox
			db.Where("processed = ?", false).Find(&outboxes)

			for _, outbox := range outboxes {
				message := kafka.Message{
					Value: []byte(toJSON(outbox)),
				}
				err := writer.WriteMessages(ctx, message)
				if err != nil {
					log.Println("Failed to write message to Kafka:", err)
					continue
				}

				fmt.Println("Message sent to Kafka:", outbox.ID)

				outbox.Processed = true
				db.Save(&outbox)
			}

			time.Sleep(10 * time.Second)
		}
	}
}

func toJSON(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
