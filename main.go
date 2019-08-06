package main

import (
	_ "bee/blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"), 30)
	log.Println("connect mysql success")

	orm.RunSyncdb("default", false, true)
	orm.Debug = true // print sql
	beego.SetLogger("file", `{"filename":"logs/blog.log"}`)

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetStaticPath("static", "static")

	beego.Run()
}
