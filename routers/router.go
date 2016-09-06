package routers

import (
	"github.com/yunkaiyueming/MonShark/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("home/", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/index", &controllers.HomeController{}, "GET:Index")
	beego.Router("home/register", &controllers.HomeController{}, "*:Register") //如果这个地方用POST，会导致在控制器中用this.GetString()方法无法获取到数据
}
