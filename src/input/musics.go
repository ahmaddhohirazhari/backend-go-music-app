package input

import "time"

type InputMusic struct {
	Music_ID    uint      `gorm:"primaryKey" json:"Music_id"`
	Name        string    `json:"name" valid:"type(string)"`
	Album       string    `json:"album" valid:"type(string)"`
	AlbumArt    string    `json:"image"`
	Singer      string    `json:"singer" valid:"type(string)"`
	PublishDate time.Time `json:"created_at"`
}