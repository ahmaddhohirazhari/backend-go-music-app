package musics

import (
	"backend/src/database/gorm/models"
	"backend/src/helpers"
	"errors"

	"gorm.io/gorm"
)

var response helpers.Response
type music_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *music_repo {
	return &music_repo{grm}
}

func (repo *music_repo) FindAll() (*models.Musics, error) {

	var musics models.Musics

	result := repo.db.Find(&musics)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &musics, nil
}

func (repo *music_repo) FindByID(id int) (*models.Music, error) {

	var musics models.Music

	result := repo.db.First(&musics, id)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &musics, nil
}



func (repo *music_repo) Add(data *models.Music) (*models.Music, error) {

	var musics models.Music

	result := repo.db.Create(data)

if result.Error != nil {
    return nil, errors.New("gagal menambahkan data")
}


getdata := repo.db.First(&musics, &data.Music_ID)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

return &musics, nil


	
}

func (repo *music_repo) Delete(id int) (*models.Music, error) {

	var musics models.Music

	getdata := repo.db.First(&musics, id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	result := repo.db.Delete(&models.Music{}, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &musics, nil
}

func (repo *music_repo) Update(id int, data *models.Music) (*models.Music, error) {

	var musics models.Music

	result := repo.db.Model(&models.Music{}).Where("music_id = ?", id).Updates(&models.Music{Name: data.Name, })

	if result.Error != nil {
		return nil, errors.New("gagal meng-update data")
	}

	getData := repo.db.First(&musics, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &musics, nil
}

