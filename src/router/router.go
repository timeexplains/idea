package  router

import "github.com/astaxie/beego"
import "controller"

func init() {
	beego.Router("/", &controller.MainController{})
}