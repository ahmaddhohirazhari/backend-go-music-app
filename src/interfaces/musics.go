package interfaces

import (
	"backend/src/database/gorm/models"
	"backend/src/helpers"
	"backend/src/input"
	"mime/multipart"
)

type MusicRepo interface {
	FindAll() (*models.Musics, error)
	FindByID(id int) (*models.Music, error)
	Add(data *models.Music) (*models.Music, error)
	Delete(id int) (*models.Music, error)
	Update(id int, data *models.Music) (*models.Music, error)
}

type MusicService interface {
	FindAll() (*helpers.Response, error)
	FindByID(id int) (*helpers.Response, error)
	Add(data *input.InputMusic, file multipart.File, handle *multipart.FileHeader) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	Update(id int, data *models.Music) (*helpers.Response, error)
}
