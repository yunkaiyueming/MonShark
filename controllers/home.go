package controllers

import (
	_ "encoding/json"
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
	this.getMachineConfig()
	this.MyRender("home/view_machine.html")
}

//获取机器的性能，Mgo的各项参数等
func (this *HomeController) getMachineConfig() {
	machineConfigData := map[string]MachineConfig{
		"bi":        {Name: "bi", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"oa":        {Name: "oa", Ip: "s119.29.00.59", User: "00", Port: 10220},
		"rsdk-set":  {Name: "rsdk-set", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"bi2-agent": {Name: "bi2-agent", Ip: "s119.00.25.59", User: "00", Port: 10220},
	}

	this.Data["machineConfigData"] = machineConfigData
}

//数据管理
func (this *HomeController) ShowMgoData() {
	db := this.GetString("db")
	if db == "" {
		db = "test"
	}

	mgoCols := this.GetColsByDb(db)
	col := this.GetString("col")
	if col == "" {
		col = mgoCols[0]
	}
	fmt.Println(col)

	mgoDbs := this.GetMgoDbs()
	docs := this.GetDocByCol(col)

	this.Data["db"] = db
	this.Data["col"] = col
	this.Data["docs"] = docs
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

func (this *HomeController) GetDocByCol(colName string) interface{} {
	var docs []interface{}

	db := this.GetString("db")
	if db == "" {
		db = "test"
	}
	err := this.mgoSession.DB(db).C(colName).Find(nil).All(&docs)
	helpers.CheckError(err)

	fmt.Println(docs)
	return docs
}

func (this *HomeController) AjaxGetColByDb() {
	db := this.GetString("db")
	cols := this.GetColsByDb(db)

	fmt.Println(cols)
	this.Data["json"] = cols
	this.ServeJSON()
}
