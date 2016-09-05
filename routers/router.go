package routers

import (
	"github.com/yunkaiyueming/MonShark/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
