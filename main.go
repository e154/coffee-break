package main

import (
    "./core"
    . "./core/capi"
	"runtime"
)

func main() {

    NewGuiApplication()
    core.Run(ApplicationThread())
	ApplicationExec()
}

func init() {
    runtime.LockOSThread()
}