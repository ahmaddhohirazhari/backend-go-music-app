package musics

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/src/database/gorm/models"
	"backend/src/helpers"
	"backend/src/input"
	"backend/src/interfaces"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type musics_ctrl struct {
	svc interfaces.MusicService
}

func NewCtrl(ctrl interfaces.MusicService) *musics_ctrl {
	return &musics_ctrl{ctrl}
}

func (ctrl *musics_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *musics_ctrl) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataID = r.URL.Query()
	id := dataID.Get("id")

	data, err := ctrl.svc.FindByID(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *musics_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")

	var data input.InputMusic
	var decoder = schema.NewDecoder()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("music")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		helpers.New(err.Error(), 400, true).Send(w)
		return
	}

	err = helpers.Validation(data.Name, data.Album, data.Singer)
	if err != nil {
		helpers.New(err.Error(), 400, true).Send(w)
		return
	}

	// _, err = govalidator.ValidateStruct(data)
	// if err != nil {
	// 	helpers.New(err.Error(), 400, true).Send(w)
	// 	return
	// }

	result, err := ctrl.svc.Add(&data, file, handler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}

func (ctrl *musics_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id := data["id"]

	result, err := ctrl.svc.Delete(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *musics_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.Music

	json.NewDecoder(r.Body).Decode(&data)

	id := dataId.Get("id")

	result, err := ctrl.svc.Update(id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}
