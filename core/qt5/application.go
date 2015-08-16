package qt5

/*

#include "application.h"

*/
import "C"

type Application struct {
    ptr C.cQApplication
}

func NewApplication(argc int, argv string) Application {
    var instance Application
    instance.ptr = C.Application(C.int(argc), C.CString(argv))
    return instance
}

func (i Application) Free() {
    C.FreeApplication(i.ptr)
}

func (i Application) Exec() {
    C.ExecApplication(i.ptr)
}