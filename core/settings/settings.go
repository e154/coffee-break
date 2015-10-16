package settings

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/user"
    "time"
	"strconv"

	"github.com/astaxie/beego/config"

	"../capi"
)

const (
    CONF_NAME string = "app.conf"
    APP_NAME string = "coffeebreak"
    APP_VERSION = "1.2.0"
    permMode os.FileMode = 0666
)

// Singleton
var instantiated *Settings = nil

func SettingsPtr() *Settings {
    if instantiated == nil {
        instantiated = new(Settings);
    }
    return instantiated;
}

type Settings struct {
    HomeDir string
    Idle time.Duration
    Work time.Duration
    LockConst time.Duration
    WorkConst time.Duration         // Таймер рабочег времени
    TotalIdle time.Duration
    TotalWork time.Duration
    Tick time.Duration
    Protect time.Duration
    StartTime time.Time
    UpTime time.Duration
    cfg config.ConfigContainer
    dir string
    Last_stage string
    Ready bool
    Paused bool
    SoundEnabled bool
    RunAtStartup bool
    Message_title string
    Message_body string
    Message_image string
    Default_timer time.Duration         // Дефолтный таймер рабочего времени
    Alarm_file string
    Webserver_address string
    SysTray capi.SystemTray
    Run_at_startup bool
    Lock time.Duration
    AppVersion string
	LockScreen int
}

func (s *Settings) GetHomeDir() (string, error) {
    
    if len(s.HomeDir) != 0 {
        return s.HomeDir, nil
    }

    user, err := user.Current()
    if err != nil {
        return s.HomeDir, err
    }

    s.HomeDir = user.HomeDir

    return s.HomeDir, nil
}

func (s *Settings) Init() *Settings {

    if len(s.HomeDir) == 0 {
        s.GetHomeDir()
    }

    s.StartTime = time.Now()
    s.dir = fmt.Sprintf("%s/.%s/", s.HomeDir, APP_NAME)

    s.Paused = false
    s.SoundEnabled = true
    s.RunAtStartup = true
    s.WorkConst = 2700 * time.Second // 45min
    s.LockConst = 600 * time.Second // 10min
    s.Tick = 1 * time.Second
    s.Protect = 30 * time.Second
    s.Message_title = "Внимание"
    s.Message_body = "Вы работаете уже {work_time}, через {time_to_lock} экран будет заблокирован, позаботьтесь о сохранности данных!\n\r"
    s.Message_image = "static_source/images/icons/watch-grey.png"
    s.Default_timer = 2700 * time.Second
    s.Alarm_file = "static_source/audio/aperture_logo_bells_01_01.wav"
    s.Webserver_address = "127.0.0.1:8078"
    s.AppVersion = APP_VERSION
    s.LockScreen = 1

//    create app conf dir
    fileList, _ := ioutil.ReadDir(s.HomeDir)

    var exist bool
    for _, file := range fileList {
        if file.Name() == "."+APP_NAME {
            exist = true
            break
        }
    }

    if !exist {
        dir := fmt.Sprintf(`%s/.%s`, s.HomeDir, APP_NAME)
        os.MkdirAll(dir, os.ModePerm)
    }

    return s
}

func (s *Settings) Save() (*Settings, error) {

    if len(s.HomeDir) == 0 {
        s.GetHomeDir()
    }

    if _, err := os.Stat(s.dir + CONF_NAME); os.IsNotExist(err) {
        ioutil.WriteFile(s.dir + CONF_NAME, []byte{}, permMode)
    }

    cfg, err := config.NewConfig("ini", s.dir + CONF_NAME)
    if err != nil {
        return s, err
    }

    cfg.Set("paused", fmt.Sprintf("%t", s.Paused))
    cfg.Set("run_at_startup", fmt.Sprintf("%t", s.RunAtStartup))
    cfg.Set("sound_enabled", fmt.Sprintf("%t", s.SoundEnabled))
    cfg.Set("lock", fmt.Sprintf("%v", s.LockConst.Seconds()))
    cfg.Set("work", fmt.Sprintf("%v", s.WorkConst.Seconds()))
    cfg.Set("protect", fmt.Sprintf("%v", s.Protect.Seconds()))
    cfg.Set("message_title", s.Message_title)
    cfg.Set("message_body", s.Message_body)
    cfg.Set("message_image", s.Message_image)
    cfg.Set("default_timer", fmt.Sprintf("%v", s.Default_timer.Seconds()))
    cfg.Set("alarm_file", s.Alarm_file)
    cfg.Set("webserver_address", s.Webserver_address)
    cfg.Set("app_version", s.AppVersion)
    cfg.Set("lock_screen", strconv.Itoa(s.LockScreen))

    if err := cfg.SaveConfigFile(s.dir + CONF_NAME); err != nil {
        fmt.Printf("err with create conf file: %s\n", s.dir + CONF_NAME)
        return s, err
    }

    return s, nil
}

func (s *Settings) Load() (*Settings, error) {

//    fmt.Printf("read config: %s\n", s.dir + CONF_NAME)

    if _, err := os.Stat(s.dir + CONF_NAME); os.IsNotExist(err) {
        return s.Save()
    }

    // read config file
    cfg, err := config.NewConfig("ini", s.dir + CONF_NAME)
    if err != nil {
        return s, err
    }

    second := func(key string) time.Duration {
        val, _ := cfg.Int(key)
        return time.Duration(val) * time.Second
    }


    if cfg.String("app_version") != APP_VERSION {
        s.Save()
        return s.Load()
    }

    s.Ready = true
    s.SoundEnabled, _ = cfg.Bool("sound_enabled")
    s.RunAtStartup, _ = cfg.Bool("run_at_startup")
    s.Paused, _ = cfg.Bool("paused")
    s.LockConst = second("lock")
    s.WorkConst = second("work")
    s.Protect = second("protect")
    s.Message_title = cfg.String("message_title")
    s.Message_body = cfg.String("message_body")
    s.Message_image = cfg.String("message_image")
    s.Default_timer = second("default_timer")
    s.Alarm_file = cfg.String("alarm_file")
    s.Webserver_address = cfg.String("webserver_address")
    s.AppVersion = cfg.String("app_version")
    s.LockScreen, _ = cfg.Int("lock_screen")

    return s, nil
}
