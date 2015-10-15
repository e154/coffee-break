package notify

import (
    "fmt"
    "os"
    "time"
    "path/filepath"
)

const DELAY = 10000;

func Show(title, body, image string) {

    Init("message")

    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    msg := NotificationNew(title,body, dir + "/" + image)

    if msg == nil {
        fmt.Fprintf(os.Stderr, "Unable to create a new notification\n")
        return
    }
    NotificationSetTimeout(msg, DELAY)

    // msg.Show()
    if e := NotificationShow(msg); e != nil {
        fmt.Fprintf(os.Stderr, "%s\n", e.Message())
        return
    }

    time.Sleep(DELAY * time.Microsecond)
    
    // msg.Close()
    NotificationClose(msg)

    UnInit()
}
