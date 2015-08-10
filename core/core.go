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

var (
    display *C.Display
    isWork bool
    signal_count int
    last_idle_time time.Duration
    tmp_idle_timer time.Duration
    settings *Settings
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
    last_idle_time = settings.Idle
    settings.Idle = time.Duration(*idle) * time.Millisecond
}

func fsm() {

    isWork = settings.Idle < time.Second
    protected := settings.Idle < settings.Protect

    if settings.Paused {
        settings.Work += settings.Tick
        return
    }

    switch settings.Stage {
        case "work":
            settings.Work += settings.Tick
            settings.TotalWork += settings.Tick
            if isWork {
                if settings.Work > settings.WorkConst {
                    send_signal("work_idle")
                }

            } else {

                if !protected {
                    settings.Stage = "idle"
                }
            }
        case "idle":
            tmp_idle_timer += settings.Tick
            settings.TotalIdle += settings.Tick
            if !isWork {
                if settings.Idle > settings.IdleConst {
                    send_signal("idle_work")
                }

            } else {
                if tmp_idle_timer < settings.IdleConst {
                    send_signal("unfinished_idle")
                } else {
                    tmp_idle_timer = 0
                    settings.Work = 0
                    settings.Stage = "work"
                }
            }
    }

    fmt.Printf("\n")
    fmt.Printf("settings.Idle: %v\n", settings.Idle)
    fmt.Printf("last_idle_time: %v\n", last_idle_time)
    fmt.Printf("PROTECT_INTERVAR: %v\n", settings.Protect)
    fmt.Printf("protected: %t\n", protected)
    fmt.Printf("settings.Stage: %s\n", settings.Stage)
    fmt.Printf("isWork: %t\n", isWork)
    fmt.Printf("IdleConst: %v\n", settings.IdleConst)
    fmt.Printf("WorkConst: %v\n", settings.WorkConst)
}

func strConverter(in string) (out string) {

    out = strings.Replace(in, "{idle_time}", fmt.Sprintf("%v", settings.Idle), -1)
    out = strings.Replace(out, "{work_time}", fmt.Sprintf("%v", settings.Work), -1)
    out = strings.Replace(out, "{idle}", fmt.Sprintf("%v", settings.IdleConst), -1)
    return
}

func send_signal(stage string) {

    if signal_count > 10 {
        signal_count = 0
    } else {
        signal_count++
        return
    }

    switch stage {
        case "idle_work":
            go notify.Show(strConverter(settings.Idle_work_title), strConverter(settings.Idle_work_body), strConverter(settings.Idle_work_image))

        case "work_idle":
            go notify.Show(strConverter(settings.Work_idle_title), strConverter(settings.Work_idle_body), strConverter(settings.Work_idle_image))

        case "unfinished_idle":
            go notify.Show(strConverter(settings.Unfinished_idle_title), strConverter(settings.Unfinished_idle_body), strConverter(settings.Unfinished_idle_image))
    }
}

func Run() {

    // init settings
    settings = SettingsPtr()
    settings.Init()
    settings.Load()

    // set start time
    fmt.Printf("running ...\n")
    fmt.Printf("current time %s\n", settings.StartTime)

    // timer
    go func() {
        ticker := time.Tick(settings.Tick)
        for {
            select {
            case <-ticker:
                settings.UpTime = time.Now().Sub(settings.StartTime)
                go idle_time()
                go fsm()
            }
        }
    }()

    webserver.Run()

    // node
    node.Run()
}
