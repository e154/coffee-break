package qt5

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -Lcore/qt5 -L/home/delta54/Qt5.5.0/5.5/gcc/lib -lQt5Gui -lQt5Core -lQt5Widgets -lstdc++  -lsystray

#include "systemtray.h"

*/
import "C"

type SystemTray struct {
    ptr C.cSystemTray
}

func NewSystemTray(s string) SystemTray {
    var instance SystemTray
    instance.ptr = C.Init(C.CString(s))
    return instance
}

func (i SystemTray) Free() {
    C.Free(i.ptr)
}

func (i SystemTray) SetTrayIcon(img string) {
    C.SetTrayIcon(i.ptr, C.CString(img))
}

func (i SystemTray) SetTrayToolTip(title string) {
    C.SetTrayToolTip(i.ptr, C.CString(title))
}