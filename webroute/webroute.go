package webroute

import (
	"file-sever/controller"

	"github.com/gorilla/mux"
)

func SetRoute(r *mux.Router) {
	r.HandleFunc("/api/siss/upload/image", controller.Upload).Methods("POST")
	r.HandleFunc("/api/siss/extract/image/{imageName}", controller.Extract).Methods("GET")
}
