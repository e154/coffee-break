package capi

/*
#cgo CXXFLAGS: -I.
#cgo LDFLAGS: -L. -Lcore/capi -lQt5Core -lQt5Gui -lQt5Widgets -lstdc++  -lcapi

#include "capi.h"

*/
import "C"

import (
    "unsafe"
)

/*--------------------------------------------------------
/* Application
--------------------------------------------------------*/
func NewGuiApplication() { C.NewGuiApplication() }
func ApplicationExec() { C.ApplicationExec() }
func ApplicationExit() { C.ApplicationExit() }
func ApplicationFlushAll() { C.ApplicationFlushAll() }
func ApplicationPtr() unsafe.Pointer { return C.ApplicationPtr() }

/*--------------------------------------------------------
/* SystemTray
--------------------------------------------------------*/
type SystemTray struct {
    addr unsafe.Pointer
}

func GetSystemTray() SystemTray {
    var instance SystemTray
    instance.addr = C.GetSystemTray()
    return instance
}

func (t SystemTray) SetIcon(img string) {
    C.SetTrayIcon((t.addr), C.CString(img))
}

func (t SystemTray) SetToolTip(tooltip string) {
    C.SetTrayToolTip(t.addr, C.CString(tooltip))
}

func (t SystemTray) SetVisible(b bool) {
    C.SetTrayVisible(t.addr, C._Bool(b))
}

func (t SystemTray) SetTimeCallback(callback unsafe.Pointer) {
    C.SetTimeCallback(t.addr, callback)
}

func (t SystemTray) SetTime(time int) {
    C.SetTime((t.addr), C.int(time))
}

func (t SystemTray) GetTime() int {
    return int(C.GetTime(t.addr))
}

//export go_callback_int
func go_callback_int(pfoo unsafe.Pointer, p1 C.int) {
    foo := *(*func(C.int))(pfoo)
    foo(p1)
}