package api

import (
    "net/http"
    "github.com/gorilla/mux"
)

func StartServer() {
    r := mux.NewRouter()
    // Define routes
    // e.g., r.HandleFunc("/api/upload", uploadFileHandler).Methods("POST")
    http.ListenAndServe(":8080", r)
}
