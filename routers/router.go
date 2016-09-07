package routers

import (
	"github.com/astaxie/beego"
	"github.com/yunkaiyueming/MonShark/controllers"
)

func init() {
	//数据管理模块
	beego.Router("/", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/index", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/ShowMgoData", &controllers.HomeController{}, "GET:ShowMgoData")

	//用户模块
	beego.Router("user/", &controllers.UserController{}, "GET:InsertTest")
}
