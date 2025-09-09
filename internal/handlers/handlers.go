package handlers

import "net/http"


func MainHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodGet {
        http.Error(w, "Сервер не поддерживает "+req.Method, http.StatusMethodNotAllowed)
        return
    }
    http.ServeFile(w, req, "../index.html")
}