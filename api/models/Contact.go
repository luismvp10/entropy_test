package models

import (
	//"errors"
	//"html"
	//"strings"
	//"time"

	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	//"time"
)

type Contact struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nombre    string `gorm:"size:255;not null;" json:"name"`
	Apodo     string `gorm:"size:255;not null;" json:"nickname"`
	Email     string `gorm:"size:100;not null;" json:"email"`
	NumeroTel string `gorm:"size:10;not null;" json:"telephone_number"`
	Direccion string `gorm:"size:255;not null;" json:"address"`
	User      User   `json:"user"`
	UserID    uint32 `gorm:"not null" json:"user_id"`
}

func (c *Contact) Prepare() {
	c.ID = 0
	c.Nombre = html.EscapeString(strings.TrimSpace(c.Nombre))
	c.Apodo = html.EscapeString(strings.TrimSpace(c.Apodo))
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.NumeroTel = html.EscapeString(strings.TrimSpace(c.NumeroTel))
	c.Direccion = html.EscapeString(strings.TrimSpace(c.Direccion))
	c.User = User{}

}

func (c *Contact) Validate() error {

	if c.Nombre == "" {
		return errors.New("Required Name")
	}
	if c.Apodo == "" {
		return errors.New("Required Nickname")
	}
	if c.Email == "" {
		return errors.New("Required Email")
	}

	if c.NumeroTel == "" {
		return errors.New("Required Telephone Number")
	}
	if c.Direccion == "" {
		return errors.New("Required Address")
	}
	if c.UserID < 1 {
		return errors.New("Required User ID")
	}
	return nil
}

func (c *Contact) SaveContact(db *gorm.DB) (*Contact, error) {
	var err error
	err = db.Debug().Model(&Contact{}).Create(&c).Error
	if err != nil {
		return &Contact{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.User).Error
		if err != nil {
			return &Contact{}, err
		}
	}
	return c, nil
}

func (c *Contact) FindContactByUserId(db *gorm.DB, pid uint64) (*[]Contact, error) {
	var err error
	contacts := []Contact{}

	err = db.Debug().Model(&Contact{}).Where("user_id = ?", pid).Find(&contacts).Limit(100).Error
	if err != nil {
		return &[]Contact{}, err
	}

	if len(contacts) > 0 {
		for i, _ := range contacts {
			err := db.Debug().Model(&User{}).Where("id = ?", contacts[i].UserID).Take(&contacts[i].User).Error
			if err != nil {
				return &[]Contact{}, err
			}
		}
	}

	return &contacts, nil
}

func (c *Contact) FindContact(db *gorm.DB, pid uint64) (*Contact, error) {
	var err error
	err = db.Debug().Model(&Contact{}).Where("id = ?", pid).Take(&c).Error
	if err != nil {
		return &Contact{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.User).Error
		if err != nil {
			return &Contact{}, err
		}
	}
	return c, nil
}
