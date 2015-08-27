package core

import (
    "C"
    "fmt"
    "unsafe"
    "time"

    "github.com/looplab/fsm"
    st "./settings"
    api "./capi"
    "./audio"
    "./webserver"
    "./xprint"
)

const (
    DoubleClick int = 2
    SingleClick int = 3
)

var (
    isWork bool
    watcher *Watcher
    settings *st.Settings
    systray api.SystemTray
    player *audio.Player
)

type Watcher struct {
    FSM *fsm.FSM
}

func (w *Watcher) enterPause(e *fsm.Event) {
    systray.SetIcon("static_source/images/icons/watch-grey.png")
    settings.Paused = true
}

func (w *Watcher) leavePause(e *fsm.Event) {
    systray.SetIcon("static_source/images/icons/watch-blue.png")
    settings.Paused = false
}

func (w *Watcher) enterWork(e *fsm.Event) {

}

//func (w *Watcher) enterWork(e *fsm.Event) {}

func (w *Watcher) enterState(e *fsm.Event) {
    fmt.Printf("Enter state %s\n", e.Dst)
}

func Run() {

    // init settings
    settings = st.SettingsPtr()
    settings.Init()
    settings.Load()

    systrayInit()
    playerInit()
    loopInit()
    webserverInit()

    // watcher init
    watcher = new(Watcher)

    watcher.FSM = fsm.NewFSM(
        "paused",
        fsm.Events{
            // Рабочее состояние, до момента "Х" более 5 минут
            {Name: "work", Src: []string{"paused"}, Dst: "worked"},

            // Рабочее стостояние, до момента "Х" менее 5 минут
            {Name: "work_lock", Src: []string{"worked"}, Dst: "work_locked"},

            // Рабочее стостояние, до момента "Х" менее 1 минут
            {Name: "work_warning_lock", Src: []string{"work_locked"}, Dst: "work_warning_locked"},

            // Момент "Х"
            {Name: "lock", Src: []string{"work_warning_locked"}, Dst: "locked"},

            // Пауза, все процессы остановлены
            {Name: "pause", Src: []string{"worked"}, Dst: "paused"},
        },
        fsm.Callbacks{
            "enter_paused": func(e *fsm.Event) { watcher.enterPause(e) },
            "leave_paused": func(e *fsm.Event) { watcher.leavePause(e) },
            "enter_state": func(e *fsm.Event) { watcher.enterState(e) },
            "enter_work": func(e *fsm.Event) { watcher.enterWork(e) },
        },
    )

    err := watcher.FSM.Event("work")
    if err != nil {
        fmt.Println(err)
    }

    err = watcher.FSM.Event("pause")
    if err != nil {
        fmt.Println(err)
    }
}

func loop() {
    isWork = settings.Idle < time.Second
//    protected := settings.Idle < settings.Protect

    if settings.Paused {
        settings.Work += settings.Tick
        return
    }

    switch watcher.FSM.Current() {
        case "worked":
            settings.Work += settings.Tick
            settings.TotalWork += settings.Tick


        case "paused":
            //
    }
}

func systrayInit() {

    seconds := func(d time.Duration) int {
        ns := d.Nanoseconds()
        return int(ns / 1000000000)
    }

    systray = api.GetSystemTray()
    systray.SetIcon("static_source/images/icons/watch-red.png")
    systray.SetToolTip("Watcher")

    var TimeCallbackFunc = TimeCallback
    var DTimeCallbackFunc = DTimeCallback
    var IconActivatedCallbackFunc = IconActivatedCallback
    var RunAtStartupCallbackFunc = RunAtStartupCallback

    systray.SetTimeCallback(unsafe.Pointer(&TimeCallbackFunc))
    systray.SetDTimeCallback(unsafe.Pointer(&DTimeCallbackFunc))
    systray.SetIconActivatedCallback(unsafe.Pointer(&IconActivatedCallbackFunc))
    systray.SetRunAtStartupCallback(unsafe.Pointer(&RunAtStartupCallbackFunc))

    systray.SetVisible(true)

    // set value
    if settings != nil && settings.Default_timer != 0 {
        systray.SetDTime( seconds(settings.Default_timer) )
    }

}

func playerInit() {

    player = audio.PlayerPtr()
    if settings.Alarm_file != "" {
        player.File("static_source/audio/" + settings.Alarm_file)
    }
}

func loopInit() {

    go func() {
        ticker := time.Tick(settings.Tick)
        for {
            select {
            case <-ticker:
                settings.UpTime = time.Now().Sub(settings.StartTime)
                go xprint.Update()
                loop()
            }
        }
    }()
}

func webserverInit() {
    webserver.Run(settings.Webserver_address)
}

// systray callbacks
func TimeCallback(x C.int) {

}

func DTimeCallback(x C.int) {

    settings.Default_timer = time.Duration(x) * time.Second
    settings.Save()
}

func IconActivatedCallback(x C.int) {

    switch int(x) {
        case DoubleClick:
            if watcher.FSM.Current() != "paused" {
                watcher.FSM.Event("pause")
            } else {
                watcher.FSM.Event("work")
            }

        case SingleClick:

    }
}

func RunAtStartupCallback(x C.int) { fmt.Println("run at startup callback", x) }
