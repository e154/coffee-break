package main

import (
//    "log"
//    daemon "github.com/tyranron/daemonigo"
    "./core"
    "./core/qt5"
    st "./core/settings"
    "os"
)

func main() {
    // Daemonizing echo server application.
//    switch isDaemon, err := daemon.Daemonize(); {
//        case !isDaemon:
//        return
//        case err != nil:
//        log.Fatalf("main(): could not start daemon, reason -> %s", err.Error())
//    }
    // From now we are running in daemon process.

    app := qt5.NewApplication(len(os.Args), os.Args[0])

    core.Run()

    settings := st.SettingsPtr()
    settings.SysTray = qt5.NewSystemTray("Hello from stdio\n")
    settings.SysTray.SetTrayIcon("static_source/images/icons/watch-blue.png")
    settings.SysTray.SetTrayToolTip("Watcher")

    app.Exec()
}
