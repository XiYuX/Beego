package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1311795565@qq.com"
	c.TplName = "index.tpl"
}
