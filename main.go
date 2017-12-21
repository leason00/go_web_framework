package main

import (
	_ "apiproject/routers"
	_ "apiproject/models"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "apiproject/database/myredis"
	_ "apiproject/database/mymysql"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
