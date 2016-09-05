package routers

import (
	"github.com/yunkaiyueming/MonShark/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	
	beego.Router("home/", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/index", &controllers.HomeController{}, "GET:Index")
}
