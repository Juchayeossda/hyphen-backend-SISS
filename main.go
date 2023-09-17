package main

import (
	"file-sever/webroute"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8083"

func main() {
	router := mux.NewRouter()
	webroute.SetRoute(router)
	log.Println(PORT, "is started")
	http.ListenAndServe(PORT, router)
}
