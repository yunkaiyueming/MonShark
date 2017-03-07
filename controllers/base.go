package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/yunkaiyueming/MonShark/helpers"
	"github.com/yunkaiyueming/MonShark/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseController struct {
	beego.Controller

	layoutFile  string
	headerFile  string
	sidebarFile string
	footerFile  string

	mgoSession *mgo.Session
}

type WebAna struct {
	Id_        bson.ObjectId `bson:"_id"`
	Url        string        `bson:"url"`
	PageView   int           `bson:"page_view"`
	Refer      string        `bson:"refer"`
	CreateTime string        `bson:"create_time"`
	LastTime   string        `bson:"last_time"`
}

const DefaultMgoDbName = "MonShark"

func (this *BaseController) Prepare() {
	fmt.Println("base controller prepare")

	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"

	this.ConnMongoDB()
	this.RecordPageView()
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

func (this *BaseController) MyRedirect(baseUrl string, code int) {
	realUrl := helpers.SiteUrl(baseUrl)
	this.Redirect(realUrl, code)
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
		this.MyRedirect("user/login", 302)
		return false
	}
}

func (this *BaseController) GetSessionUser() interface{} {
	return this.GetSession("email")
}

func (this *BaseController) ConnMongoDB() {
	url := beego.AppConfig.String("mongoUrl")
	this.mgoSession = models.GetDbConn(url)
}

func (this *BaseController) CloseMongoDB() {
	this.mgoSession.Close()
}

func (this *BaseController) GetMgoDbs() []string {
	dbs, err := this.mgoSession.DatabaseNames()
	helpers.CheckError(err)
	return dbs
}

func (this *BaseController) GetColsByDb(dbName string) []string {
	cols, err := this.mgoSession.DB(dbName).CollectionNames()
	helpers.CheckError(err)
	return cols
}

func (this *BaseController) JsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) AjaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg

	this.JsonResult(out)
}

//web请求统计
func (this *BaseController) RecordPageView() {
	webCol := this.mgoSession.DB(DefaultMgoDbName).C("web_analystics")
	controllerName, actionName := this.GetControllerAndAction()
	url := controllerName + "/" + actionName
	refer := this.Ctx.Request.Referer()

	result := WebAna{}
	err := webCol.Find(bson.M{"url": url}).One(&result)
	helpers.CheckError(err)

	//upsert
	if result.Id_ != "" {
		fmt.Println(result)
		err := webCol.UpdateId(result.Id_, &bson.M{"$set": bson.M{"page_view": (result.PageView + 1), "last_time": helpers.MyNowDate()}})
		helpers.CheckError(err)
	} else {
		data := &WebAna{bson.NewObjectId(), url, 1, refer, helpers.MyNowDate(), helpers.MyNowDate()}
		err := webCol.Insert(data)
		helpers.CheckError(err)
	}

}
