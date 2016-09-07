package controllers

import (
	"fmt"
	_ "strings"

	"github.com/yunkaiyueming/MonShark/helpers"
)

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

//数据管理
func (this *HomeController) ShowMgoData() {
	db := this.GetString("db")
	this.GetString("col")
	if db == "" {
		db = "test"
	}

	mgoDbs := this.GetMgoDbs()
	fmt.Println(mgoDbs)
	mgoCols := this.GetColsByDb(db)

	this.Data["mgoDbs"] = mgoDbs
	this.Data["mgoCols"] = mgoCols
	this.MyRender("home/view_showMgoData.html")
}

func (this *HomeController) GetMgoDbs() []string {
	dbs, err := this.mgoSession.DatabaseNames()
	helpers.CheckError(err)
	return dbs
}

func (this *HomeController) GetColsByDb(dbName string) []string {
	cols, err := this.mgoSession.DB(dbName).CollectionNames()
	helpers.CheckError(err)
	return cols
}

func (this *HomeController) GetDocByCol(colName string) {

}
