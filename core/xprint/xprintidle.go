package xprint

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -Lcore/xprint/ -lxprintidle -lX11 -lXss -lXdmcp -lXext
#include "xprintidle.h"
*/
import "C"

import (
    "fmt"
    "time"
    st "../settings"
)

var (
    settings *st.Settings
    display *C.Display
    last_idle_time time.Duration
)

func Update() {

    if settings == nil {
        return
    }

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

func init() {
    settings = st.SettingsPtr()
}