package webserver

import (
    "net/http"
    "html/template"
    "time"
    "fmt"
    "github.com/gorilla/websocket"
)

const (
// Time allowed to write a message to the client.
// Время разрешено написать сообщение клиенту.
    writeWait = 10 * time.Second

// Time allowed to read the next message from the client.
// Время разрешено читать следующее сообщение от клиента.
    readWait = 60 * time.Second

// Отправить пинги к клиенту с этим периодом.
// Должна быть меньше, чем readWait.
    pingPeriod = (readWait * 9) / 10

// Максимальный размер сообщений от клиента.
    maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // разрешаем подключение к сокету с разных доменов
    CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

    ws, err := upgrader.Upgrade(w, r, nil)
    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400)
        return
    } else if err != nil {
        fmt.Printf(err.Error())
        return
    }

    // регистрация нового клиента
    c := &Client{
        Send: make(chan []byte, 512),
        Ws: ws,
    }

    H.Register <- c
    go c.WritePump()
    c.ReadPump()
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

func Run(address string) {

    // routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/ws", wsHandler)
    http.HandleFunc("/js/", fileHandler)
    http.HandleFunc("/css/", fileHandler)
    http.HandleFunc("/images/", fileHandler)
    http.HandleFunc("/templates/", fileHandler)

    go http.ListenAndServe(address, nil)
}
