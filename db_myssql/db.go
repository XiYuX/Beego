package db_myssql

import (
	"BeegoPackage0922/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//在初始化函数中连接数据库
func init() {
	fmt.Println("连接数据库")
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接失败，请重试")
	}
	Db = db


}

func InsertUser(user models.NewPerson)(int64,error){
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名：",user.Nick,user.Brithday,user.Address,user.Password)
	result,err:=Db.Exec("insert into user(nick,Brithday,Address,password) values(?,?,?,?)",user.Nick,user.Brithday,user.Address,user.Password)
	if err != nil{
		return -1,err
	}
	id ,err := result.RowsAffected()
	if err != nil{
		return -1,err
	}
	return id,nil
}

