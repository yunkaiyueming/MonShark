package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/yunkaiyueming/MonShark/models"
	"gopkg.in/mgo.v2"
)

type BaseController struct {
	beego.Controller

	layoutFile  string
	headerFile  string
	sidebarFile string
	footerFile  string

	mgoSession *mgo.Session
}

func (this *BaseController) Prepare() {
	fmt.Println("base controller prepare")

	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"

	this.ConnMongoDB()
}

func (this *BaseController) MyRender(viewFile string) {
	this.Layout = this.layoutFile
	this.TplName = viewFile

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headerFile"] = this.headerFile
	this.LayoutSections["footerFile"] = this.footerFile
	this.LayoutSections["sidebarFile"] = this.sidebarFile

	this.PrepareViewData()
	this.Render()
}

func (this *BaseController) PrepareViewData() {
	staticUrl := beego.AppConfig.String("static_url")
	siteUrl := beego.AppConfig.String("siteUrl")

	this.Data["staticUrl"] = staticUrl
	this.Data["siteUrl"] = siteUrl
}

func (this *BaseController) CheckLogin() bool {
	email := this.GetSession("email")
	if email != nil {
		return true
	} else {
		return false
	}
}

func (this *BaseController) ConnMongoDB() {
	url := beego.AppConfig.String("mongoUrl")
	this.mgoSession = models.GetDbConn(url)
}

func (this *BaseController) CloseMongoDB() {
	this.mgoSession.Close()
}

func (this *BaseController) LoginRender(viewFile string) {
	this.TplName = viewFile
	this.PrepareViewData()
	this.Render()
}

func (this *BaseController) getMachineConfig() {
	machineConfigData := map[string]MachineConfig{
		"bi":        {Name: "bi", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"oa":        {Name: "oa", Ip: "s119.29.00.59", User: "00", Port: 10220},
		"rsdk-set":  {Name: "rsdk-set", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"bi2-agent": {Name: "bi2-agent", Ip: "s119.00.25.59", User: "00", Port: 10220},
	}

	this.Data["machineConfigData"] = machineConfigData
}
