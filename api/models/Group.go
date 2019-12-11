package models

import (
	"time"
)

type Group struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nombre    string    `gorm:"size:255;not null;" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
