package core

import (
    "C"
    "fmt"
    "unsafe"
    "time"
    "strings"

    "github.com/looplab/fsm"
    st "./settings"
    api "./capi"
    "./audio"
    "./notify"
    "./webserver"
    "./xprint"
)

const (
    DoubleClick = 2
    SingleClick = 3
)

var (
    isWork bool
    watcher *Watcher
    settings *st.Settings
    systray api.SystemTray
    player *audio.Player
    window api.MainWindow
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
    settings.Work = 0
}

func (w *Watcher) enterWork(e *fsm.Event) {
    settings.Work = 0
    systray.SetIcon("static_source/images/icons/watch-blue.png")
}

func (w *Watcher) enterWorkLock(e *fsm.Event) {
    systray.SetIcon("static_source/images/icons/watch-red.png")
    showNotify()
}

func (w *Watcher) enterWorkWarningLock(e *fsm.Event) {
    systray.SetIcon("static_source/images/icons/watch-red.png")
    showNotify()
}

func (w *Watcher) enterState(e *fsm.Event) {
    fmt.Printf("Enter state %s\n", e.Dst)
}

func (w *Watcher) enterLock(e *fsm.Event) {
    window.FullScreen()
    systray.SetIcon("static_source/images/icons/watch-red.png")
}

func (w *Watcher) leaveLock(e *fsm.Event) {
    window.Hidde()
    settings.Lock = 0
    settings.Work = 0
}

func Run() {

    // init settings
    settings = st.SettingsPtr()
    settings.Init()
    settings.Load()

    systrayInit()
    playerInit()
    webserverInit()
    windowInit()
    fsmInit()
    loopInit()
}

func loop() {
    isWork = settings.Idle < time.Second
//    protected := settings.Idle < settings.Protect

    if settings.Paused {
        return
    }

    settings.Work += settings.Tick
    settings.TotalWork += settings.Tick

    switch watcher.FSM.Current() {
        case "worked":
            if isWork {
                if settings.Work >= (settings.WorkConst - 5 * time.Minute) {
                    watcher.FSM.Event("work_lock")
                } else if settings.Work >= (settings.WorkConst - 1 * time.Minute) {
                    watcher.FSM.Event("work_warning_lock")
                }
            }

        case "work_locked":
            if settings.Work < (settings.WorkConst - 5 * time.Minute) {
                watcher.FSM.Event("work")
            } else if settings.Work >= (settings.WorkConst - 1 * time.Minute) {
                watcher.FSM.Event("work_warning_lock")
            }

        case "work_warning_locked":
            if settings.Work <= (settings.WorkConst - 1 * time.Minute) {
                watcher.FSM.Event("work_locked")
            } else if settings.Work >= (settings.WorkConst) {
                watcher.FSM.Event("lock")
            }

        case "paused":


        case "locked":
            settings.Lock += settings.Tick
            settings.TotalIdle += settings.Tick

            if settings.Lock >= settings.LockConst {
                watcher.FSM.Event("work")
            }

    }
}

func systrayInit() {

    seconds := func(d time.Duration) int {
        ns := d.Nanoseconds()
        return int(ns / 1000000000)
    }

    thread := api.ApplicationThread()
    systray = api.GetSystemTray()
    systray.MoveToThread(thread)
    systray.SetIcon("static_source/images/icons/watch-grey.png")
    systray.SetToolTip("Watcher")

    var TimeCallbackFunc = TimeCallback
    var DTimeCallbackFunc = DTimeCallback
    var IconActivatedCallbackFunc = IconActivatedCallback
    var RunAtStartupCallbackFunc = RunAtStartupCallback
    var AlarmCallbackFunc = AlarmCallback

    systray.SetTimeCallback(unsafe.Pointer(&TimeCallbackFunc))
    systray.SetDTimeCallback(unsafe.Pointer(&DTimeCallbackFunc))
    systray.SetIconActivatedCallback(unsafe.Pointer(&IconActivatedCallbackFunc))
    systray.SetRunAtStartupCallback(unsafe.Pointer(&RunAtStartupCallbackFunc))
    systray.SetAlarmCallback(unsafe.Pointer(&AlarmCallbackFunc))

    systray.SetVisible(true)

    // set value
    if settings != nil && settings.Default_timer != 0 {
        systray.SetDTime( seconds(settings.Default_timer) )
        systray.SetTime( seconds(settings.Default_timer) )
    }

    if settings.RunAtStartup {
        systray.SetRunAtStartup(1)
    } else {
        systray.SetRunAtStartup(0)
    }

    if settings.SoundEnabled {
        systray.SetAlarm(1)
        systray.SetAlarmInfo("Alarm is on")
    } else {
        systray.SetAlarm(3)
        systray.SetAlarmInfo("Alarm is off")
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
    settings.WorkConst = time.Duration(x) * time.Second
    settings.Work = 0
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

func RunAtStartupCallback(x C.int) {
    if int(x) == 1 {
        settings.RunAtStartup = true
    } else {
        settings.RunAtStartup = false
    }

    settings.Save()
}

func AlarmCallback(x C.int) {

    if int(x) == 1 {
        settings.SoundEnabled = true
        systray.SetAlarmInfo("Alarm is on")
    } else {
        settings.SoundEnabled = false
        systray.SetAlarmInfo("Alarm is off")
    }

    settings.Save()
}

func strConverter(in string) (out string) {

    out = strings.Replace(in, "{idle_time}", fmt.Sprintf("%v", settings.Idle), -1)
    out = strings.Replace(out, "{work_time}", fmt.Sprintf("%v", settings.Work), -1)
    out = strings.Replace(out, "{lock}", fmt.Sprintf("%v", settings.LockConst), -1)
    out = strings.Replace(out, "{time_to_lock}", fmt.Sprintf("%v", settings.WorkConst - settings.Work), -1)
    return
}

func showNotify() {

    if settings.Work <= ( 3 * time.Minute) {
        systray.ShowMessage(strConverter(settings.Message_title), strConverter(settings.Message_body), 1)
        return
    }

    if settings.SoundEnabled {
        player.Play()
    }

    go notify.Show(strConverter(settings.Message_title), strConverter(settings.Message_body), settings.Message_image)
}

func fsmInit() {

    watcher = new(Watcher)

    watcher.FSM = fsm.NewFSM(
    "paused",
    fsm.Events{
        // Рабочее состояние, до момента "Х" более 5 минут
        {Name: "work", Src: []string{"paused", "work_locked", "work_warning_locked", "locked"}, Dst: "worked"},

        // Рабочее состояние, до момента "Х" менее 5 минут
        {Name: "work_lock", Src: []string{"worked", "locked"}, Dst: "work_locked"},

        // Рабочее состояние, до момента "Х" менее 1 минут
        {Name: "work_warning_lock", Src: []string{"work_locked", "locked"}, Dst: "work_warning_locked"},

        // Момент "Х"
        {Name: "lock", Src: []string{"work_warning_locked"}, Dst: "locked"},

        // Пауза, все процессы остановлены
        {Name: "pause", Src: []string{"worked", "work_locked", "locked", "work_warning_locked"}, Dst: "paused"},
    },
    fsm.Callbacks{
        "enter_paused": func(e *fsm.Event) { watcher.enterPause(e) },
        "leave_paused": func(e *fsm.Event) { watcher.leavePause(e) },
        "enter_state": func(e *fsm.Event) { watcher.enterState(e) },
        "enter_worked": func(e *fsm.Event) { watcher.enterWork(e) },
        "enter_work_locked": func(e *fsm.Event) { watcher.enterWorkLock(e) },
        "enter_work_warning_locked": func(e *fsm.Event) { watcher.enterWorkWarningLock(e) },
        "enter_locked": func(e *fsm.Event) { watcher.enterLock(e) },
        "leave_locked": func(e *fsm.Event) { watcher.leaveLock(e) },
    },
    )

    if settings.RunAtStartup {
        err := watcher.FSM.Event("work")
        if err != nil {
            fmt.Println(err)
        }
    }
}

func windowInit() {

    window = api.GetMainWindow()
    window.Url(fmt.Sprintf("http://%s/lock", settings.Webserver_address))
}