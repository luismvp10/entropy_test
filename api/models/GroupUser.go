package models

import ()

type GroupUser struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	User    User   `json:"user"`
	UserID  uint32 `gorm:"not null" json:"user_id"`
	Group   Group  `json:"group"`
	GroupID uint32 `gorm:"not null" json:"group_id"`
	Admin   bool   `gorm:"not null" json:"admin_status"`
}
