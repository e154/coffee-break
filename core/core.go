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
    LOCK_ICON = "static_source/images/icons/watch-red.png"
    WORK_ICON = "static_source/images/icons/watch-blue.png"
    PAUSE_ICON = "static_source/images/icons/watch-grey.png"
)

var (
    isWork bool
    watcher *Watcher
    settings *st.Settings
    systray api.SystemTray
    player *audio.Player
    window api.MainWindow
    TimeCallbackFunc = TimeCallback
	DTimeCallbackFunc = DTimeCallback
	IconActivatedCallbackFunc = IconActivatedCallback
	RunAtStartupCallbackFunc = RunAtStartupCallback
	AlarmCallbackFunc = AlarmCallback
	LockScreenCallbackFunc = LockScreenCallback
)

type Watcher struct {
    FSM *fsm.FSM
}

func (w *Watcher) enterPause(e *fsm.Event) {
    systray.SetIcon(PAUSE_ICON)
    settings.Paused = true
}

func (w *Watcher) leavePause(e *fsm.Event) {
    systray.SetIcon(WORK_ICON)
    settings.Paused = false
    settings.Work = 0
}

func (w *Watcher) enterWork(e *fsm.Event) {
    settings.Work = 0
    systray.SetIcon(WORK_ICON)
}

func (w *Watcher) enterWorkLock(e *fsm.Event) {
    systray.SetIcon(LOCK_ICON)
    showNotify()
}

func (w *Watcher) enterWorkWarningLock(e *fsm.Event) {
    systray.SetIcon(LOCK_ICON)
    showNotify()
}

func (w *Watcher) enterState(e *fsm.Event) {
    fmt.Printf("Enter state %s\n", e.Dst)
}

func (w *Watcher) enterLock(e *fsm.Event) {
	windowUrl()
    window.FullScreen()
    systray.SetIcon(LOCK_ICON)
}

func (w *Watcher) leaveLock(e *fsm.Event) {
    window.Hidde()
    settings.Lock = 0
    settings.Work = 0
}

func Run(thread unsafe.Pointer) {

    // init settings
    settings = st.SettingsPtr()
    settings.Init()
    settings.Load()

    systrayInit(thread)
    playerInit()
    webserverInit()
    windowInit(thread)
    fsmInit()
    loopInit()
}

func loop() {
    isWork = settings.Idle < time.Second
//    protected := settings.Idle < settings.Protect

    if settings.Paused {
        return
    }

    if watcher.FSM.Current() != "locked" {
		if isWork {
			settings.Work += settings.Tick
			settings.TotalWork += settings.Tick
		} else {
			settings.TotalIdle += settings.Tick
		}
    }

    switch watcher.FSM.Current() {
        case "worked":
            if isWork {
                if settings.Work >= (settings.WorkConst - 5 * time.Minute) {
                    err := watcher.FSM.Event("work_lock")
                    errHandler(err)
                } else if settings.Work >= (settings.WorkConst - 1 * time.Minute) {
                    err := watcher.FSM.Event("work_warning_lock")
                    errHandler(err)
                }
            }

        case "work_locked":
            if settings.Work < (settings.WorkConst - 5 * time.Minute) {
                err := watcher.FSM.Event("work")
                errHandler(err)
            } else if settings.Work >= (settings.WorkConst - 1 * time.Minute) {
                err := watcher.FSM.Event("work_warning_lock")
                errHandler(err)
            }

        case "work_warning_locked":
            if settings.Work <= (settings.WorkConst - 1 * time.Minute) {
                err := watcher.FSM.Event("work_locked")
                errHandler(err)
            } else if settings.Work >= (settings.WorkConst) {
                err := watcher.FSM.Event("lock")
                errHandler(err)
            }

        case "paused":


        case "locked":
            settings.Lock += settings.Tick
            settings.TotalIdle += settings.Tick

            if settings.Lock >= settings.LockConst {
                err := watcher.FSM.Event("work")
                errHandler(err)
            }

    }

//    fmt.Printf("\n")
//    fmt.Printf("settings.Idle: %v\n", settings.Idle)
//    fmt.Printf("PROTECT_INTERVAR: %v\n", settings.Protect)
//    fmt.Printf("settings.Protect: %v\n", settings.Protect)
//    fmt.Printf("Stage: %s\n", watcher.FSM.Current())
//    fmt.Printf("isWork: %t\n", isWork)
//    fmt.Printf("Work: %v\n", settings.Work)
//    fmt.Printf("TotalIdle: %v\n", settings.TotalIdle)
//    fmt.Printf("LockConst: %v\n", settings.LockConst)
//    fmt.Printf("WorkConst: %v\n", settings.WorkConst)
}

func systrayInit(thread unsafe.Pointer) {

    seconds := func(d time.Duration) int {
        ns := d.Nanoseconds()
        return int(ns / 1000000000)
    }

    systray = api.GetSystemTray()
    systray.MoveToThread(thread)
    systray.SetIcon(PAUSE_ICON)
    systray.SetToolTip("Coffee Break")

    systray.SetTimeCallback(unsafe.Pointer(&TimeCallbackFunc))
    systray.SetDTimeCallback(unsafe.Pointer(&DTimeCallbackFunc))
    systray.SetIconActivatedCallback(unsafe.Pointer(&IconActivatedCallbackFunc))
    systray.SetRunAtStartupCallback(unsafe.Pointer(&RunAtStartupCallbackFunc))
    systray.SetAlarmCallback(unsafe.Pointer(&AlarmCallbackFunc))
    systray.SetLockScreenCallback(unsafe.Pointer(&LockScreenCallbackFunc))

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

	if settings.LockScreen < 1 {
		settings.LockScreen = 1
	}

	systray.SetLockScreen(settings.LockScreen)
}

func playerInit() {

    player = audio.PlayerPtr()
    if settings.Alarm_file != "" {
        player.File(settings.Alarm_file)
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
                err := watcher.FSM.Event("pause")
                errHandler(err)
            } else {
                err := watcher.FSM.Event("work")
                errHandler(err)
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

func LockScreenCallback(x C.int) {
	settings.LockScreen = int(x)
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
//        systray.ShowMessage(strConverter(settings.Message_title), strConverter(settings.Message_body), 1)
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
        errHandler(err)
    }
}

func windowInit(thread unsafe.Pointer) {

    window = api.GetMainWindow()
	window.Thread(thread)
	windowUrl()
}

func windowUrl() {

	var lock string
	switch settings.LockScreen {
	case 1:
		lock = "lockmatrix"
	case 2:
		lock = "lockbsod"
	case 3:
		lock = "lockide"
	default:
		lock = "lockmatrix"
	}
	window.Url(fmt.Sprintf("http://%s/%s", settings.Webserver_address, lock))
}

func errHandler(err error) {
    if err == nil { return }
    fmt.Printf("error: %s\n", err.Error())
}