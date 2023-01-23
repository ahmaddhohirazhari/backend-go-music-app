package models

import (
	"time"
)

type Music struct {
	Music_ID  uint      `gorm:"primaryKey" json:"Music_id"`
	Name        string    `json:"name" valid:"type(string)"`
	Album       string    `json:"album" valid:"type(string)"`
	AlbumArt    string    `json:"album_art"`
	Singer      string    `json:"singer" valid:"type(string)"`
	PublishDate time.Time `json:"publish_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Musics []Music
