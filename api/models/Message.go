package models

import (
	"time"
)

type Message struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Message    string    `gorm:"size:255;not null;" json:"messsage"`
	UserTo     User      `json:"user_to"`
	UsertToID  uint32    `gorm:"null"json:"userTo_id"`
	User       User      `json:"user_from"`
	UserFromID uint32    `gorm:"not null" json:"userFrom_id"`
	Group      Group     `json:"group"`
	GroupID    uint32    `gorm:"not null"json:"group_id"`
	Date       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
