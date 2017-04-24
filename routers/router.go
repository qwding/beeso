package routers

import (
	"beeso/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dynamic/*", &controllers.DynamicController{},"*:Router")
	beego.Router("/dynamicregistor", &controllers.RegistorController{})
}
