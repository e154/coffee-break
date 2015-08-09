package core

import (
    "github.com/astaxie/beego"
)

func Init() {
    beego.AppConfig.String("db_user")
}