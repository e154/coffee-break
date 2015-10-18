package capi

/*
#cgo CXXFLAGS: -I.
#cgo LDFLAGS: -L. -Lcore/capi -lQt5Core -lQt5Gui -lQt5Widgets -lQt5WebKitWidgets -lstdc++  -lcapi

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
func ApplicationThread() unsafe.Pointer { return C.ApplicationThread() }

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

func (t SystemTray) SetIcon(img string) { C.SetTrayIcon((t.addr), C.CString(img)) }
func (t SystemTray) SetToolTip(tooltip string) { C.SetTrayToolTip(t.addr, C.CString(tooltip)) }
func (t SystemTray) SetVisible(b bool) { C.SetTrayVisible(t.addr, C._Bool(b)) }
func (t SystemTray) SetTimeCallback(callback unsafe.Pointer) { C.SetTimeCallback(t.addr, callback) }
func (t SystemTray) SetTime(time int) { C.SetTime((t.addr), C.int(time)) }
func (t SystemTray) GetTime() int { return int(C.GetTime(t.addr)) }
func (t SystemTray) SetDTimeCallback(callback unsafe.Pointer) { C.SetDTimeCallback(t.addr, callback) }
func (t SystemTray) SetDTime(time int) { C.SetDTime((t.addr), C.int(time)) }
func (t SystemTray) GetDTime() int { return int(C.GetDTime(t.addr)) }
func (t SystemTray) SetAlarmCallback(callback unsafe.Pointer) { C.SetAlarmCallback(t.addr, callback) }
func (t SystemTray) SetAlarm(state int) { C.SetAlarm((t.addr), C.int(state)) }
func (t SystemTray) GetAlarm() int { return int(C.GetAlarm(t.addr)) }
func (t SystemTray) SetRunAtStartupCallback(callback unsafe.Pointer) { C.SetRunAtStartupCallback(t.addr, callback) }
func (t SystemTray) SetRunAtStartup(state int) { C.SetRunAtStartup((t.addr), C.int(state)) }
func (t SystemTray) GetRunAtStartup() int { return int(C.GetRunAtStartup(t.addr)) }
func (t SystemTray) SetAlarmInfo(info string) { C.SetAlarmInfo((t.addr), C.CString(info)) }
func (t SystemTray) GetAlarmInfo() string { return C.GoString(C.GetAlarmInfo(t.addr)) }
func (t SystemTray) SetIconActivatedCallback(callback unsafe.Pointer) { C.SetIconActivatedCallback(t.addr, callback) }
func (t SystemTray) MoveToThread(thread unsafe.Pointer) { C.MoveToThread(t.addr, thread) }
func (t SystemTray) ShowMessage(title, message string, icon int) { C.ShowMessage(t.addr, C.CString(title), C.CString(message), C.int(icon)) }
func (t SystemTray) SetLockScreenCallback(callback unsafe.Pointer) { C.SetLockScreenCallback(t.addr, callback) }
func (t SystemTray) SetLockScreen(state int) { C.SetLockScreen((t.addr), C.int(state)) }
func (t SystemTray) GetLockScreen() int { return int(C.GetLockScreen(t.addr)) }

//export go_callback_int
func go_callback_int(pfoo unsafe.Pointer, p1 C.int) {
    foo := *(*func(C.int))(pfoo)
    foo(p1)
}

/*--------------------------------------------------------
/* Main window
--------------------------------------------------------*/
type MainWindow struct {
    addr unsafe.Pointer
}

func GetMainWindow() *MainWindow {
    instance := new(MainWindow)
    instance.addr = C.GetMainWindow()
    return instance
}

func (w *MainWindow) Show() { C.MainWindowShow(w.addr) }
func (w *MainWindow) Hidde() { C.MainWindowHidde(w.addr) }
func (w *MainWindow) ShowNormal() { C.MainWindowNormal(w.addr) }
func (w *MainWindow) FullScreen() { C.MainWindowFullScreen(w.addr) }
func (w *MainWindow) Url(url string) { C.MainWindowUrl(w.addr, C.CString(url)) }
func (w *MainWindow) Thread(thread unsafe.Pointer) { C.MainWindowThread(w.addr, thread) }
func (w *MainWindow) Delete() { C.MainWindowDelete(w.addr) }
