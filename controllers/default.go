package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "http://www.baidu.com"
	c.Data["Email"] = "2585717835@qq.com"
	c.TplName = "index.tpl"
}
