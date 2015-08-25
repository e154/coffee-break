package systemtray

import (
    "C"
    . "../capi"
    "fmt"
    "unsafe"
)

func TimeCallback(x C.int) { fmt.Println("time callback", x) }
func DTimeCallback(x C.int) { fmt.Println("dtime callback", x) }
func IconActivatedCallback(x C.int) { fmt.Println("icon callback", x) }
func RunAtStartupCallback(x C.int) { fmt.Println("run at startup callback", x) }

func Run() {

    NewGuiApplication()

    systray := GetSystemTray()
    systray.SetIcon("static_source/images/icons/watch-red.png")
    systray.SetToolTip("Watcher")
    systray.SetVisible(true)

    var TimeCallbackFunc = TimeCallback
    var DTimeCallbackFunc = DTimeCallback
    var IconActivatedCallbackFunc = IconActivatedCallback
    var RunAtStartupCallbackFunc = RunAtStartupCallback

    systray.SetTimeCallback(unsafe.Pointer(&TimeCallbackFunc))
    systray.SetDTimeCallback(unsafe.Pointer(&DTimeCallbackFunc))
    systray.SetIconActivatedCallback(unsafe.Pointer(&IconActivatedCallbackFunc))
    systray.SetRunAtStartupCallback(unsafe.Pointer(&RunAtStartupCallbackFunc))

    systray.SetTime(45 * 60)
    systray.SetDTime(45 * 60)
    systray.SetAlarm(2)
    systray.SetAlarmInfo("test msg")

    ApplicationExec()
}

