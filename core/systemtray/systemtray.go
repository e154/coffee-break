package systemtray

import (
    . "../capi"
//    "fmt"
    "fmt"
)

func Run() {

    NewGuiApplication()

    // systray
    dialog := NewQDialog(nil)
    systray := NewSystemTray(dialog.Ptr())
    systray.SetIcon("static_source/images/icons/watch-blue.png")
    systray.SetToolTip("Watcher")


    // actions
    quitAction := NewQAction("/icon", "&Quit", ApplicationPtr())

    // connectors
    quitAction.On("hovered", func(){
        fmt.Println("obj got a click")
        systray.SetIcon("static_source/images/icons/watch-red.png")
    })

    // menu
    trayIconMenu := NewQMenu(dialog.Ptr())
    trayIconMenu.Clear()
    trayIconMenu.AddAction(quitAction)

    systray.SetContextMenu(trayIconMenu)
    systray.SetVisible(true)


    ApplicationExec()
}