package main

import (
	_ "BeegoPackage0922/db_myssql"
	_ "BeegoPackage0922/routers"
	"github.com/astaxie/beego"
)

func main() {
	//config := beego.AppConfig
	//appName :=config.String("appname")
	//fmt.Println("应用名称",appName)
	//port ,err :=config.Int("httpport")
	//if err != nil{
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}
	//
	//fmt.Println(port)
	beego.Run()
}

