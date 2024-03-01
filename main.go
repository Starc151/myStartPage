package main

import (
    "html/template"
    "net/http"
    "os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("index.html"))
    tpl.Execute(w, nil)
}


func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("css"))
    mux.Handle("/css/", http.StripPrefix("/css/", fs))

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}