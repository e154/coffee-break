package capi

/*
#cgo CXXFLAGS: -std=c++0x -Wall -fno-strict-aliasing -I.
#cgo LDFLAGS: -L. -Lcore/capi -L/usr/lib -lQt5Core -lQt5Gui -lQt5Widgets -lstdc++  -lcapi

#include "capi.h"

*/
import "C"

import (
    "unsafe"
    "reflect"
    "gopkg.in/qml.v1/cdata"
    "sync/atomic"
    "fmt"
)

var (
    guiFunc      = make(chan func())
    guiDone      = make(chan struct{})
    guiMainRef   uintptr
    guiPaintRef  uintptr
    guiLock      = 0
    guiIdleRun   int32
)

func NewGuiApplication() { C.NewGuiApplication() }
func ApplicationExec() { C.ApplicationExec() }
func ApplicationExit() { C.ApplicationExit() }
func ApplicationFlushAll() { C.ApplicationFlushAll() }
func ApplicationThread() unsafe.Pointer { return C.ApplicationThread() }
func ApplicationPtr() unsafe.Pointer { return C.ApplicationPtr() }

type QObject struct {
    ptr unsafe.Pointer
}

type SystemTray struct { QObject }
type QWidget struct { QObject }
type QDialog struct { QObject }
type QAction struct { QObject }
type QMenu struct { QObject }

/*--------------------------------------------------------
/* QObject
--------------------------------------------------------*/
func (o *QObject) MoveToThread(thread unsafe.Pointer) {
    C.QObjectMoveToThread(o.ptr, thread)
}

func (o *QObject) Thread() unsafe.Pointer {
    return C.QObjectThread(o.ptr)
}

func (o QObject) Ptr() unsafe.Pointer {
    return o.ptr
}

/*--------------------------------------------------------
/* SystemTray
--------------------------------------------------------*/
func NewSystemTray(parent unsafe.Pointer) SystemTray {
    var instance SystemTray
    instance.ptr = C.NewSystemTray( parent )
    return instance
}

func (t SystemTray) SetIcon(img string) { C.SetTrayIcon(t.ptr, C.CString(img)) }
func (t SystemTray) SetToolTip(tooltip string) { C.SetTrayToolTip(t.ptr, C.CString(tooltip)) }
func (t SystemTray) SetVisible(b bool) { C.SetTrayVisible(t.ptr, C._Bool(b)) }
func (t SystemTray) SetContextMenu(menu QMenu) { C.SetTrayContextMenu(t.ptr, menu.Ptr()) }

/*--------------------------------------------------------
/* QWidget
--------------------------------------------------------*/
func NewQWidget(parent unsafe.Pointer) QWidget {
    var instance QWidget
    instance.ptr = C.NewQDialog( parent )
    return instance
}

/*--------------------------------------------------------
/* QDialog
--------------------------------------------------------*/
func NewQDialog(parent unsafe.Pointer) QDialog {
    var instance QDialog
    instance.ptr = C.NewQDialog( parent )
    return instance
}

/*--------------------------------------------------------
/* QAction
--------------------------------------------------------*/
func NewQAction(icon, text string, parent unsafe.Pointer) QAction {
    var instance QAction
    instance.ptr = C.NewQAction( C.CString(icon), C.CString(text), parent )
    return instance
}
func (t QAction) SetIcon(img string) { C.QActionSetIcon(t.ptr, C.CString(img)) }
func (t QAction) Icon() string { return C.GoString(C.QActionIcon(t.ptr)) }

/*--------------------------------------------------------
/* QMenu
--------------------------------------------------------*/
func NewQMenu(parent unsafe.Pointer) QMenu {
    var instance QMenu
    instance.ptr = C.NewQmenu( parent )
    return instance
}
func (m QMenu) AddAction(action QAction) { C.QMenuAddAction(m.ptr, action.Ptr()) }
func (m QMenu) RemoveAction(action QAction) { C.QMenuRemoveAction(m.ptr, action.Ptr()) }
func (m QMenu) InsertAction(before, action QAction) { C.QMenuInsertAction(m.ptr, before.Ptr(), action.Ptr()) }
func (m QMenu) Clear() { C.QMenuClear(m.ptr) }

var connectedFunction = make(map[*interface{}]bool)

// On connects the named signal from obj with the provided function, so that
// when obj next emits that signal, the function is called with the parameters
// the signal carries.
//
// The provided function must accept a number of parameters that is equal to
// or less than the number of parameters provided by the signal, and the
// resepctive parameter types must match exactly or be conversible according
// to normal Go rules.
//
// For example:
//
//     obj.On("clicked", func() { fmt.Println("obj got a click") })
//
// Note that Go uses the real signal name, rather than the one used when
// defining QML signal handlers ("clicked" rather than "onClicked").
//
// For more details regarding signals and QML see:
//
//     http://qt-project.org/doc/qt-5.0/qtqml/qml-qtquick2-connections.html
//
func (obj *QObject) On(signal string, function interface{}) {
    fmt.Println("On enter")
    funcv := reflect.ValueOf(function)
    funct := funcv.Type()
    if funcv.Kind() != reflect.Func {
        panic("function provided to On is not a function or method")
    }
    if funct.NumIn() > C.MaxParams {
        panic("function takes too many arguments")
    }
    csignal, csignallen := unsafeStringData(signal)
    var cerr *C.error
    RunMain(func() {
        cerr = C.objectConnect(obj.Ptr(), csignal, csignallen, ApplicationPtr(), unsafe.Pointer(&function), C.int(funcv.Type().NumIn()))
        fmt.Println("stage 4...")
        if cerr == nil {
            fmt.Println("stage 10...")
            connectedFunction[&function] = true
//            stats.connectionsAlive(+1)
            return
        }
    })
    error(cerr)
}

func error(cerr *C.error) {
    fmt.Printf("error: %s", C.GoString((*C.char)(unsafe.Pointer(cerr))))
}

// RunMain runs f in the main QML thread and waits for f to return.
//
// This is meant to be used by extensions that integrate directly with the
// underlying QML logic.
func RunMain(f func()) {
    fmt.Println("Run main")
    ref := cdata.Ref()
    if ref == guiMainRef || ref == atomic.LoadUintptr(&guiPaintRef) {
        fmt.Println("stage 6...")
        // Already within the GUI or render threads. Attempting to wait would deadlock.
        f()
        return
    }

    fmt.Println("stage 7...")
    // Tell Qt we're waiting for the idle hook to be called.
//    if atomic.AddInt32(&guiIdleRun, 1) == 1 {
//        C.IdleTimerStart()
//    }

    // Send f to be executed by the idle hook in the main GUI thread.
    guiFunc <- f

    fmt.Println("stage 8...")
    // Wait until f is done executing.
    <-guiDone

    fmt.Println("stage 9...")
}

func init() {
    var f func()

    go func(){
        for {
            select {
            case f = <-guiFunc:
            default:

            }

            f()

        }
    }()
}

type Post struct {
    Title    string `json:"title"`
    Views    int    `json:"pageviews"`
    Password string `json:"-"`
}