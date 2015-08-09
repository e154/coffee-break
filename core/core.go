package core

/*
#cgo CFLAGS: -Ixprint
#cgo LDFLAGS: -Lxprint/ -Lcore/xprint/ -lxprintidle -lX11 -lXss -lXdmcp -lXext
#include "xprintidle.h"
*/
import "C"

import (
    "fmt"
    "time"
    "./notify"
    "strings"
    "./node"
    "./webserver"
)

func up_time() {
    // set up time
    UP_TIME = time.Now().Sub(START_TIME)
}

var (
    START_TIME time.Time
    UP_TIME time.Duration
    IDLE_TIME time.Duration
    LAST_IDLE_TIME time.Duration
    WORK_TIME time.Duration
    IDLE_CONST time.Duration = 15 * time.Minute
    WORK_CONST time.Duration = 45 * time.Minute
    PROTECT_INTERVAR time.Duration = 30 * time.Minute
    ONE_TICK time.Duration = 1 * time.Second
    STAGE string = "work" // work|idle|signal
    HOMEDIR string
    IDLE_WORK_TITLE string
    IDLE_WORK_BODY string
    IDLE_WORK_IMAGE string
    WORK_IDLE_TITLE string
    WORK_IDLE_BODY string
    WORK_IDLE_IMAGE string
    UNFINISHED_IDLE_TITLE string
    UNFINISHED_IDLE_BODY string
    UNFINISHED_IDLE_IMAGE string
    display *C.Display
    PAUSE bool
    isWork bool
    SIGNAL_COUNT int
    TOTAL_IDLE time.Duration
    TOTAL_WORK time.Duration
    TMP_IDLE_TIMER time.Duration
    PORT int = 8080
)

func idle_time() {

    if display == nil {
        display = C.getDisplay()
        if display == nil {
            fmt.Printf("error: couldn't open display\n")
            return
        }
    }

    idle := new(C.ulong)
    err := C.getIdle(idle, display)
    if err != 0 {

        var err_text string
        switch err{
            case 2:
            err_text = "screen saver extension not supported"
            case 3:
            err_text = "couldn't query screen saver info"
            default:
            err_text = "unknow"
        }

        fmt.Printf("error: %s\n", err_text)
        return
    }
    LAST_IDLE_TIME = IDLE_TIME
    IDLE_TIME = time.Duration(*idle) * time.Millisecond
}

func fsm() {

    isWork = IDLE_TIME < time.Second
    protected := IDLE_TIME < PROTECT_INTERVAR

    if PAUSE {
        WORK_TIME += ONE_TICK
        return
    }

    switch STAGE {
        case "work":
            WORK_TIME += ONE_TICK
            TOTAL_WORK += ONE_TICK
            if isWork {
                if WORK_TIME > WORK_CONST {
                    send_signal("work_idle")
                }

            } else {

                if !protected {
                    STAGE = "idle"
                }
            }
        case "idle":
            TMP_IDLE_TIMER += ONE_TICK
            TOTAL_IDLE += ONE_TICK
            if !isWork {
                if IDLE_TIME > IDLE_CONST {
                    send_signal("idle_work")
                }

            } else {
                if TMP_IDLE_TIMER < IDLE_CONST {
                    send_signal("unfinished_idle")
                } else {
                    TMP_IDLE_TIMER = 0
                    WORK_TIME = 0
                    STAGE = "work"
                }
            }
    }

//    fmt.Printf("\n")
//    fmt.Printf("IDLE_TIME: %v\n", IDLE_TIME)
//    fmt.Printf("LAST_IDLE_TIME: %v\n", LAST_IDLE_TIME)
//    fmt.Printf("PROTECT_INTERVAR: %v\n", PROTECT_INTERVAR)
//    fmt.Printf("protected: %t\n", protected)
//    fmt.Printf("STAGE: %s\n", STAGE)
//    fmt.Printf("isWork: %t\n", isWork)
}

func strConverter(in string) (out string) {

    out = strings.Replace(in, "{idle_time}", fmt.Sprintf("%v", IDLE_TIME), -1)
    out = strings.Replace(out, "{work_time}", fmt.Sprintf("%v", WORK_TIME), -1)
    out = strings.Replace(out, "{idle_const}", fmt.Sprintf("%v", IDLE_CONST), -1)
    return
}

func send_signal(stage string) {

    if SIGNAL_COUNT > 10 {
        SIGNAL_COUNT = 0
    } else {
        SIGNAL_COUNT++
        return
    }

    switch stage {
        case "idle_work":
            go notify.Show(strConverter(IDLE_WORK_TITLE), strConverter(IDLE_WORK_BODY), IDLE_WORK_IMAGE)

        case "work_idle":
            go notify.Show(strConverter(WORK_IDLE_TITLE), strConverter(WORK_IDLE_BODY), WORK_IDLE_IMAGE)

        case "unfinished_idle":
            go notify.Show(strConverter(UNFINISHED_IDLE_TITLE), strConverter(UNFINISHED_IDLE_BODY), UNFINISHED_IDLE_IMAGE)
    }
}

func Run() {

    // set start time
    START_TIME = time.Now()
    fmt.Printf("running ...\n")
    fmt.Printf("current time %s\n", START_TIME)

    // timer
    go func() {
        ticker := time.Tick(ONE_TICK)
        for {
            select {
            case <-ticker:
                go up_time()
                go idle_time()
                go fsm()
            }
        }
    }()

    webserver.Run()

    // node
    node.Run()
}
