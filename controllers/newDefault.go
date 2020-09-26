package controllers

import (
	"BeegoPackage0922/db_myssql"
	"BeegoPackage0922/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type NewController struct {
	beego.Controller
}
//该方法用于处理post请求
func (c *NewController) Post() {
	//接收json请求
	var person models.NewPerson
	dataByte, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	//进行数据校验
	err = json.Unmarshal(dataByte, &person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("姓名：", person.Name)
	fmt.Println("生日：", person.Brithday)
	fmt.Println("住址：", person.Address)
	fmt.Println("昵称：", person.Nick)
	c.Ctx.WriteString("数据解析成功")

	id, err := db_myssql.InsertUser(person)
	if err != nil {
		c.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)
	c.Ctx.WriteString("用户保存成功")


	result := models.User{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	c.Data["json"] = &result
	c.ServeJSON()
}