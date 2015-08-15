package main

import (
//    "log"
//    daemon "github.com/tyranron/daemonigo"
    "./core"
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

    core.Run()
}
