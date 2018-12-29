package routers

import (
	"liteblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.ErrorController(&controllers.ErrorController{})
    beego.Include(&controllers.IndexController{})
}
