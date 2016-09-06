package routers

import (
	"github.com/astaxie/beego"
	"github.com/yunkaiyueming/MonShark/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("home/", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/index", &controllers.HomeController{}, "GET:Index")

	beego.Router("user/", &controllers.UserController{}, "GET:InsertTest")
}
