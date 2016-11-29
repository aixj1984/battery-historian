package main

import (
	_ "routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbdsn := beego.AppConfig.String("demo::dbdsn")
	dbdriver := beego.AppConfig.String("demo::dbdriver")
	orm.RegisterDataBase("default", dbdsn, dbdriver)

}

func main() {
	beego.SetStaticPath("/swagger", "swagger")
	orm.Debug = true
	beego.Run()
}
