package input

import "time"

type InputMusic struct {
	Music_ID    string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" valid:"type(string)"`
	Album       string    `json:"album" valid:"type(string)"`
	AlbumArt    string    `json:"image"`
	Singer      string    `json:"singer" valid:"type(string)"`
	PublishDate time.Time `json:"created_at"`
}
