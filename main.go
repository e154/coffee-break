package main

import (
    "./core"
    . "./core/capi"
)

func main() {

    NewGuiApplication()

    core.Run()

    ApplicationExec()
}
