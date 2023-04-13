package musics

import (
	"backend/src/database/gorm/models"
	"backend/src/helpers"
	"backend/src/input"
	"backend/src/interfaces"
	"mime/multipart"
)

type music_service struct {
	repo interfaces.MusicRepo
}

func NewService(svc interfaces.MusicRepo) *music_service {
	return &music_service{svc}
}

func (svc *music_service) FindAll() (*helpers.Response, error) {

	result, err := svc.repo.FindAll()
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *music_service) FindByID(id string) (*helpers.Response, error) {

	result, err := svc.repo.FindByID(id)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

// func (svc *music_service) Save(data *models.Music) (*helpers.Response, error) {

// 	var musics models.Music

// 	_, err := govalidator.ValidateStruct(data)
// 	if err != nil {
// 		res := response.ResponseJSON(400, musics)
// 		res.Message = err.Error()
// 		return res, nil
// 	}

// 	result, err := svc.repo.Add(data)
// 	if err != nil {
// 		res := response.ResponseJSON(400, result)
// 		res.Message = err.Error()
// 		return res, nil
// 	}

// 	res := response.ResponseJSON(200, result)
// 	return res, nil
// }

func (svc *music_service) Delete(id string) (*helpers.Response, error) {

	result, err := svc.repo.Delete(id)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *music_service) Update(id string, data *models.Music) (*helpers.Response, error) {
	result, err := svc.repo.Update(id, data)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *music_service) Add(data *input.InputMusic, file multipart.File, handle *multipart.FileHeader) (*helpers.Response, error) {

	var music models.Music

	images, err := helpers.UploadImages("music-app", file, handle)
	if err != nil {
		res := helpers.New("err.Error()", 400, true)
		return res, nil
	}

	music.Name = data.Name
	music.Album = data.Album
	music.AlbumArt = images.URL
	music.Singer = data.Singer
	music.PublishDate = data.PublishDate

	result, err := svc.repo.Add(&music)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}
