package controllers

import (
	"fmt"
	"taskmanage/utils"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	m1 := make(map[string]interface{})
	m1["int"] = 123
	m1["string"] = "hello"
	m1["time"] = time.Now().Format("2006-01-02 15:04:05")

	var ti interface{}

	// ti = m1

	// err := utils.StringSet("dede", ti)

	// fmt.Println(err)
	err1 := utils.StringGet("dede", &ti)

	fmt.Println(ti)
	fmt.Println(err1)

	m := make(map[string]interface{})
	m["int"] = 123
	m["string"] = "hello"
	m["bool"] = m1
	m["ceshi"] = ti
	c.Data["json"] = m
	c.ServeJSON()
}
