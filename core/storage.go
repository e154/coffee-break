package core

import (
    "os/user"
    "io/ioutil"
    "os"
    "fmt"
    "time"

    "github.com/astaxie/beego/config"
)

const (
    CONF_NAME string = "app.conf"
    PROG_NAME string = ".45minut"
    permMode os.FileMode = 0666
)

func check_home_dir() {

    if len(HOMEDIR) == 0 {
        return
    }

    // get dir list in home
    fileList, err := ioutil.ReadDir(HOMEDIR)
    if err != nil {
        return
    }

    var exist bool
    for _, file := range fileList {
        if file.Name() == PROG_NAME {
            exist = true
            break
        }
    }

    if !exist {
        dir := fmt.Sprintf(`%s/%s`, HOMEDIR, PROG_NAME)
        fmt.Printf("create dir: %s\n", dir)
        os.MkdirAll(dir, os.ModePerm)
    }
}

func check_config() {

    if len(HOMEDIR) == 0 {
        return
    }

    dir := fmt.Sprintf("%s/%s/", HOMEDIR, PROG_NAME)

    // get dir list in home
    fileList, err := ioutil.ReadDir(dir)
    if err != nil {
        return
    }

    var exist bool
    for _, file := range fileList {
        if file.Name() == CONF_NAME {
            exist = true
            break
        }
    }

    file := `#app config
port = 8888
idle = 60
work = 60
protect = 30
idle_work_title = Внимание
idle_work_body = Ты отдыхаешь уже {idle_time} пора приниматся за работу!
idle_work_image =
work_idle_title = Внимание
work_idle_body = Ты работешь уже {work_time}, иди отдохни, выпей чаю!
work_idle_image =
unfinished_idle_title = Внимание
unfinished_idle_body = {idle_const} ещё не прошло, иди отдохни, выпей чаю!
unfinished_idle_image =
`

    if !exist {
        fmt.Printf("create default config file: %s\n", dir + CONF_NAME)
        ioutil.WriteFile(dir + CONF_NAME, []byte(file), permMode)
    }
}

func read_config() {

    dir := fmt.Sprintf("%s/%s/", HOMEDIR, PROG_NAME)

    // read config file
    fmt.Printf("read config: %s\n", dir + CONF_NAME)
    cnf, err := config.NewConfig("ini", dir + CONF_NAME)

    // set idle time
    idle, err := cnf.Int("idle")
    if err == nil {
        IDLE_CONST = time.Duration(idle) * time.Second
    }

    // set work time
    work, err := cnf.Int("work")
    if err == nil {
        WORK_CONST = time.Duration(work) * time.Second
    }

    // set protect time
    protect, err := cnf.Int("protect")
    if err == nil {
        PROTECT_INTERVAR = time.Duration(protect) * time.Second
    }

    IDLE_WORK_TITLE = cnf.String("idle_work_title")
    IDLE_WORK_BODY = cnf.String("idle_work_body")
    IDLE_WORK_IMAGE = cnf.String("idle_work_image")
    WORK_IDLE_TITLE = cnf.String("work_idle_title")
    WORK_IDLE_BODY = cnf.String("work_idle_body")
    WORK_IDLE_IMAGE = cnf.String("work_idle_image")
    UNFINISHED_IDLE_TITLE = cnf.String("unfinished_idle_title")
    UNFINISHED_IDLE_BODY = cnf.String("unfinished_idle_body")
    UNFINISHED_IDLE_IMAGE = cnf.String("unfinished_idle_image")

    // set port
    port, err := cnf.Int("port")
    if err == nil {
        PORT = port
    }
}

func init() {
    fmt.Printf("storage init ...\n")
    user, err := user.Current()
    if err != nil {
        return
    }

    HOMEDIR = user.HomeDir
    fmt.Printf("use home: %s\n", HOMEDIR)

    check_home_dir()
    check_config()
    read_config()
}