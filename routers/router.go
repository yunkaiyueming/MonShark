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
	beego.Router("home/AjaxGetColByDb", &controllers.HomeController{}, "GET:AjaxGetColByDb")

	//用户模块
	beego.Router("user/register", &controllers.UserController{}, "*:Register") //如果这个地方用POST，会导致在控制器中用this.GetString()方法无法获取到数据
	beego.Router("user/doregister", &controllers.UserController{}, "*:DoRegister")
	beego.Router("user/login", &controllers.UserController{}, "*:Login")
	beego.Router("user/logout", &controllers.UserController{}, "*:LogOut")
}
