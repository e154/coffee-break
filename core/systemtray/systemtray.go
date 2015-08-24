package systemtray

import (
    "C"
    . "../capi"
    "fmt"
    "unsafe"
)

func TimeCallback(x C.int) {
    fmt.Println("callback with", x)
}

func Run() {

    NewGuiApplication()

    systray := GetSystemTray()
    systray.SetIcon("static_source/images/icons/watch-red.png")
    systray.SetToolTip("Watcher")
    systray.SetVisible(true)

    var TimeCallbackFunc = TimeCallback
    systray.SetTimeCallback(unsafe.Pointer(&TimeCallbackFunc))
    systray.SetTime(60 * 4 * 60)

    ApplicationExec()
}

