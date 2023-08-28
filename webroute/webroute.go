package webroute

import (
	"file-sever/controller"

	"github.com/gorilla/mux"
)

func SetRoute(r *mux.Router) {
	r.HandleFunc("/upload/image", controller.Upload).Methods("POST")
	r.HandleFunc("/extract/image/{imageName}", controller.Extract).Methods("GET")
}
