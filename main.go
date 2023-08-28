package main

import (
	"file-sever/webroute"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	webroute.SetRoute(router)
	http.ListenAndServe(":9190", router)
}

// func main() {f
// 	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
// 		imagePath := "./files/download.png"

// 		file, err := os.Open(imagePath)
// 		if err != nil {
// 			http.Error(w, "Image not found", http.StatusNotFound)
// 			return
// 		}

// 		_, err = io.Copy(w, file)

// 		if err != nil {
// 			log.Fatal(err)
// 		}ff

// 	})

// 	http.ListenAndServe(":9190", nil)
// }
