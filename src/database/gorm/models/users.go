package models

import (
	"time"
)

type User struct {
	User_ID     uint      `gorm:"primaryKey" json:"Music_id"`
	Name        string    `json:"name" valid:"type(string)"`
	Image       string    `json:"image"`
	Email       string    `json:"email" valid:"type(string)"`
	Password    string    `json:"password" valid:"type(string)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Users []User
