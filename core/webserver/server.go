package server

import (
    "net/http"
    "html/template"
    "log"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    var homeTempl = template.Must(template.ParseFiles("static_source/templates/index.html"))
    data := struct {
        Host       string
        }{r.Host}
    homeTempl.Execute(w, data)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static_source"+r.URL.Path)
}

func Run() {

    // routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/ws", wsHandler)
    http.HandleFunc("/js/", fileHandler)
    http.HandleFunc("/css/", fileHandler)
    http.HandleFunc("/images/", fileHandler)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
