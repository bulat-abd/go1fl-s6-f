package handlers

import (
	"net/http"
	"os"
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