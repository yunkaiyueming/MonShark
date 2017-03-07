package models

import (
	_ "errors"
	_ "strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/yunkaiyueming/MonShark/helpers"
)

func init() {
	GetMysqlDB("monshark")
}

func GetMysqlDB(dbName string) {
	db, err := helpers.GetMyConfig("database", dbName)
	if err != nil {
		helpers.ErrLog("get db config error")
	}

	orm.RegisterDataBase("default", "mysql", db)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func GetMongoUrl() string {
	mongoUrl, err := helpers.GetMyConfig("database", "mongo")
	if err != nil {
		helpers.ErrLog("get db config error")
	}

	return mongoUrl
}
