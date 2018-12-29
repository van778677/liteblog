package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

//约定：如果子controller存在NestPrepare()方法，就实现了该接口
type NestPreparer interface{
	NestPrepare()
}


type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare()  {
	log.Println("BaseControll")
	//判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	c.Data["Path"] = c.Ctx.Request.RequestURI
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}


}
