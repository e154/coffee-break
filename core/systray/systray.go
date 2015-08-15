package systray

import (
    "github.com/getlantern/systray"
    "io/ioutil"
)

var (
    watch_blue, watch_red []byte
)

func Run() {
    systray.Run(settings)
}

func settings() {
    watch_blue, err := ioutil.ReadFile("static_source/images/icons/watch-blue.png")
    watch_red, err = ioutil.ReadFile("static_source/images/icons/watch-red.png")
    if err != nil {
        return
    }

    systray.SetIcon(watch_blue)
    systray.SetTitle("Awesome App")
    systray.SetTooltip("Pretty awesome超级棒")
    mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

    for {
        select {
        case <- mQuit.ClickedCh:
            systray.Quit()
        }
    }
}