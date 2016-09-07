package controllers

import _ "fmt"

type HomeController struct {
	BaseController
}

type MachineConfig struct {
	Name   string
	Ip     string
	User   string
	Port   int
	Dbinfo string
}

type Users struct {
	Email    string
	Password string
}

func (this *HomeController) Index() {
	//this.Ctx.WriteString("aaa")
	flag := this.CheckLogin()
	if flag {
		this.getMachineConfig()
		this.MyRender("home/view_machine.html")
	} else {
		this.LoginRender("home/view_welcome.html")
	}

}
