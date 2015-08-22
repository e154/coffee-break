package capi

/*
#cgo CXXFLAGS: -std=c++0x -Wall -fno-strict-aliasing -I.
#cgo LDFLAGS: -L. -Lcore/capi -L/usr/lib -lQt5Core -lQt5Gui -lQt5Widgets -lstdc++  -lcapi

#include "capi.h"

*/
import "C"

import (
    "unsafe"
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
