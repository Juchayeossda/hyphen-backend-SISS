package main

import (
	"file-sever/webroute"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	webroute.SetRoute(router)
	log.Println("[8083] SISS start")
	http.ListenAndServe(":8083", router)
}
