package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type BaseController struct {
	beego.Controller

	layoutFile  string
	headerFile  string
	sidebarFile string
	footerFile  string
	session     *mgo.Session
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
	id := this.GetSession("id")
	name := this.GetSession("name")
	if id != nil && name != nil {
		return true
	} else {
		this.Redirect("home/index", 302)
		return false
	}
}

func (this *BaseController) ConnMongoDB() {
	url := beego.AppConfig.String("mongoUrl")
	session, err := mgo.Dial(url) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	this.session = session
}
