package notify

import (
    "fmt"
    "os"
    "time"
)

const DELAY = 10000;

func Show(title, body, image string) {

    Init("message")
    msg := NotificationNew(title,body,image)

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
