package musics

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/music").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/byId", ctrl.GetByID).Methods("GET")
	route.HandleFunc("/add", ctrl.AddData).Methods("POST").Methods("POST")
	route.HandleFunc("/delete/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")

}
