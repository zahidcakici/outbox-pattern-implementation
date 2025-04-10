package models

import "gorm.io/gorm"

type Outbox struct {
	gorm.Model
	EventType string
	Payload   string
	Processed bool
}
