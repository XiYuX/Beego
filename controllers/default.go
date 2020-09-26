package controllers

import (

	"BeegoPackage0922/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}


func (c *MainController) Get() {
	//1、获取请求数据
	user := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//2、做固定数据进行数据校验
	if user != "xiyu" || password != "123456"{
		//代表错误
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误"))
		return
	}
	//校验成功
	c.Ctx.ResponseWriter.Write([]byte("恭喜，校验成功"))


	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1311795565@qq.com"
	c.TplName = "index.tpl"
}


//编写一个post方法，用于处理post请求
//func (c *MainController)Post(){
//	//接收post请求
//	name:= c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(sex)
//	//进行数据校验
//	if name != "xiyu" && age != "18"{
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("数据校验成功")
//}

//post请求
func (c *MainController)Post(){
	//解析josn格式数据
	var person models.Person
	data,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("数据接收失败")
		return
	}
	err = json.Unmarshal(data,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("姓名：",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别：",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}

