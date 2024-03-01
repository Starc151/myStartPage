package main

import (
    "net/http"
    "os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<!doctype html><head><title>asd</title></head><body><button>Найти</button></body></html>"))
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}