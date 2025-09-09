package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)


func MainHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodGet {
        http.Error(w, "Сервер не поддерживает "+req.Method, http.StatusMethodNotAllowed)
        return
    }
    path := ""
	if _, err := os.Stat("index.html"); err == nil {
		path = "index.html"
	} else if _, err := os.Stat("../index.html"); err == nil {
        path = "../index.html"
	}
    http.ServeFile(w, req, path)
}


func UploadHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        http.Error(w, "Сервер не поддерживает "+req.Method, http.StatusMethodNotAllowed)
        return
    }

	err := req.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    log.Printf("UPLOAD: data received %s", data)
	
	convertedString := service.Transcode(string(data))

	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	fileName := timestamp + filepath.Ext(handler.Filename)
	outFile, err := os.Create(fileName)
	defer outFile.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = outFile.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Конвертация выполнена: " + convertedString))
}