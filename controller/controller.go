package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

const FILE_STORAGE_PATH = "./storage"
const IMAGE_STORAGE_PATH = FILE_STORAGE_PATH + "/" + "images"

var renderer *render.Render = render.New()

type JSONM map[string]string

func Upload(w http.ResponseWriter, r *http.Request) {
	// Form 20MB 제한
	r.ParseMultipartForm(10 << 20)

	// form에서 image input tag 추출
	file, handler, err := r.FormFile("image")
	if err != nil {
		renderer.JSON(w, http.StatusNotFound, JSONM{
			"error": "can't found file in from",
		})
		return
	}
	defer file.Close()

	// 파일 확장자 추출
	fileExtension := strings.TrimPrefix(filepath.Ext(handler.Filename), ".")

	// 파일이 확장자가 아니라면
	if fileExtension == "" {
		renderer.JSON(w, http.StatusBadRequest, JSONM{
			"error": "image is not in a recognized format",
		})
		return
	}

	// 파일 Path 생성
	filePath := fmt.Sprintf("%s/%s.%s", IMAGE_STORAGE_PATH, uuid.New().String(), fileExtension)

	// 파일 저장을 위해 사전 준비
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE /*쓰기 전용 파일과, 없으면 만든다.*/, 0666)
	if err != nil {
		renderer.JSON(w, http.StatusInternalServerError, JSONM{
			"error": "no open file from backend",
		})
		return
	}
	defer f.Close()

	// 저장
	if _, err = io.Copy(f, file); err != nil {
		renderer.JSON(w, http.StatusInternalServerError, JSONM{
			"error": "can't create image",
		})
		return
	}

	// Is ok~
	renderer.JSON(w, http.StatusOK, JSONM{
		"status": "success to image upload",
		"ident":  strings.TrimPrefix(filePath, IMAGE_STORAGE_PATH+"/"),
	})
}

func Extract(w http.ResponseWriter, r *http.Request) {

	pathValues := mux.Vars(r)
	imageName, isExist := pathValues["imageName"]
	if isExist == false {
		renderer.JSON(w, http.StatusNotFound, JSONM{
			"error": "can't found path value (/extract/image/{imageName})",
		})
		return
	}

	images, err := os.ReadDir(IMAGE_STORAGE_PATH)
	if err != nil {
		renderer.JSON(w, http.StatusInternalServerError, JSONM{
			"error": "can't read directory",
		})
		return
	}

	for _, image := range images {

		if imageName == image.Name() {

			openImage, err := os.Open(IMAGE_STORAGE_PATH + "/" + image.Name())
			if err != nil {
				renderer.JSON(w, http.StatusInternalServerError, JSONM{
					"error": "can't open" + imageName,
				})
				return
			}
			defer openImage.Close()

			_, err = io.Copy(w, openImage)
			if err != nil {
				renderer.JSON(w, http.StatusInternalServerError, JSONM{
					"error": "can't response image",
				})
				return
			}

		}

	}

}
