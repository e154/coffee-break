package systemtray

import (
    . "../capi"
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

    // menu
    trayIconMenu := NewQMenu(dialog.Ptr())
    trayIconMenu.Clear()
    trayIconMenu.AddAction(quitAction)

    systray.SetContextMenu(trayIconMenu)
    systray.SetVisible(true)


    ApplicationExec()
}