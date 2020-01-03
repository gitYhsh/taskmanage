package main

import (
	_ "taskmanage/routers"

	_ "taskmanage/sysConfig"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
