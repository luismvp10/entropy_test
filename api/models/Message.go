package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"log"
	"strings"
	"time"
)

type Message struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Message    string    `gorm:"size:255;not null;" json:"message"`
	UserTo     User      `json:"user_to"`
	UsertToID  uint32    `gorm:"null"json:"userTo_id"`
	User       User      `json:"user_from"`
	UserFromID uint32    `gorm:"not null" json:"userFrom_id"`
	Group      Group     `json:"group"`
	GroupID    uint32    `gorm:"not null"json:"group_id"`
	Date       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (m *Message) Prepare() {
	m.ID = 0
	m.Message = html.EscapeString(strings.TrimSpace(m.Message))
	m.User = User{}

}

func (m *Message) Validate() error {

	if m.Message == "" {
		return errors.New("Required Message")
	}

	if m.UserFromID < 1 {
		return errors.New("Required User ID")
	}
	return nil
}

func (m *Message) SaveMessage(db *gorm.DB) (*Message, error) {
	log.Println(m)
	var err error
	err = db.Debug().Model(&Message{}).Create(&m).Error
	if err != nil {
		return &Message{}, err
	}
	//if c.ID != 0 {
	//	err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.User).Error
	//	if err != nil {
	//		return &Contact{}, err
	//	}
	//}
	return m, nil
}

func (m *Message) FindMessagesByUserId(db *gorm.DB, pid uint64) (*[]Message, error) {
	var err error
	messages := []Message{}

	err = db.Debug().Model(&Message{}).Where("usert_to_id = ?", pid).Find(&messages).Limit(100).Error
	if err != nil {
		return &[]Message{}, err
	}

	//if len(messages) > 0 {
	//	for i, _ := range messages {
	//		err := db.Debug().Model(&User{}).Where("id = ?", messages[i].UsertToID).Take(&messages[i].User).Error
	//		if err != nil {
	//			return &[]Message{}, err
	//		}
	//	}
	//}

	return &messages, nil
}

func (m *Message) FindConversations(db *gorm.DB, pid uint64) (*[]Message, error) {
	var err error
	messages := []Message{}

	err = db.Debug().Model(&Message{}).Select("distinct(usert_to_id)").Where("user_from_id = ?", pid).Find(&messages).Limit(100).Error
	if err != nil {
		return &[]Message{}, err
	}

	//if len(messages) > 0 {
	//	for i, _ := range messages {
	//		err := db.Debug().Model(&User{}).Where("id = ?", messages[i].UsertToID).Take(&messages[i].User).Error
	//		if err != nil {
	//			return &[]Message{}, err
	//		}
	//	}
	//}

	return &messages, nil
}
