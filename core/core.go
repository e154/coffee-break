package core

import (
    "fmt"
    "time"
    "strings"
    "./notify"
    "./webserver"
    "./audio"
    "./xprint"
    st "./settings"
)

var (
    isWork bool
    signal_count int
    tmp_idle_timer time.Duration
    settings *st.Settings
    player *audio.Player
)

func fsm() {
    isWork = settings.Idle < time.Second
    protected := settings.Idle < settings.Protect

    if settings.Paused {
        settings.Work += settings.Tick
        return
    }

    switch settings.Stage {
        case "work":
            settings.Work += settings.Tick
            settings.TotalWork += settings.Tick
            if isWork {
                if settings.Work > settings.WorkConst {
                    send_signal("work_idle")
                }

            } else {

                if !protected {
                    settings.Last_stage = settings.Stage
                    settings.Stage = "idle"
                }
            }
        case "idle":
            tmp_idle_timer += settings.Tick
            settings.TotalIdle += settings.Tick
            if !isWork {
                if settings.Idle > settings.IdleConst {
                    send_signal("idle_work")
                }

            } else {
                if tmp_idle_timer < settings.IdleConst {
                    send_signal("unfinished_idle")
                } else {
                    tmp_idle_timer = 0
                    settings.Work = 0
                    settings.Last_stage = settings.Stage
                    settings.Stage = "work"
                }
            }
    }

//    fmt.Printf("\n")
//    fmt.Printf("settings.Idle: %v\n", settings.Idle)
//    fmt.Printf("PROTECT_INTERVAR: %v\n", settings.Protect)
//    fmt.Printf("protected: %t\n", protected)
//    fmt.Printf("settings.Protect: %v\n", settings.Protect)
//    fmt.Printf("settings.Stage: %s\n", settings.Stage)
//    fmt.Printf("isWork: %t\n", isWork)
//    fmt.Printf("Work: %v\n", settings.Work)
//    fmt.Printf("TotalIdle: %v\n", settings.TotalIdle)
//    fmt.Printf("IdleConst: %v\n", settings.IdleConst)
//    fmt.Printf("WorkConst: %v\n", settings.WorkConst)
//    fmt.Printf("Notify_count: %d\n", settings.Notify_count)
}

func strConverter(in string) (out string) {

    out = strings.Replace(in, "{idle_time}", fmt.Sprintf("%v", settings.Idle), -1)
    out = strings.Replace(out, "{work_time}", fmt.Sprintf("%v", settings.Work), -1)
    out = strings.Replace(out, "{idle}", fmt.Sprintf("%v", settings.IdleConst), -1)
    return
}

func send_signal(stage string) {

    if signal_count > 10 {
        signal_count = 0
    } else {
        signal_count++
        return
    }

    if settings.Notify_count > settings.Maximum_notify {
        if settings.Last_stage != settings.Stage {
            settings.Notify_count = 0
        }
        return
    }

    settings.Notify_count++

    switch stage {
        case "idle_work":
            go notify.Show(strConverter(settings.Idle_work_title), strConverter(settings.Idle_work_body), strConverter(settings.Idle_work_image))

        case "work_idle":
            go notify.Show(strConverter(settings.Work_idle_title), strConverter(settings.Work_idle_body), strConverter(settings.Work_idle_image))

        case "unfinished_idle":
            go notify.Show(strConverter(settings.Unfinished_idle_title), strConverter(settings.Unfinished_idle_body), strConverter(settings.Unfinished_idle_image))
    }

    if settings.SoundEnabled {
        player.Play()
    }
}

func Run() {

    // init settings
    settings = st.SettingsPtr()
    settings.Init()
    settings.Load()

    // audio
    player = audio.PlayerPtr()
    if settings.Alarm_file != "" {
        player.File("static_source/audio/" + settings.Alarm_file)
    }

    // timer
    go func() {
        ticker := time.Tick(settings.Tick)
        for {
            select {
            case <-ticker:
                settings.UpTime = time.Now().Sub(settings.StartTime)
                go xprint.Update()
                fsm()
            }
        }
    }()

    webserver.Run(settings.Webserver_address)
}
