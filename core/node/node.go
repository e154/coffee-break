package node

import (
    "github.com/lonnc/golang-nw"
    "fmt"
    "net/http"
)

const listenAddr = "localhost:4000"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, web")
}

func Run() {

    http.HandleFunc("/", handler)
    err := http.ListenAndServe(listenAddr, nil)
    if err != nil {

    }

    nodeWebkit, _ := nw.New()
    nodeWebkit.ListenAndServe(handler)
}