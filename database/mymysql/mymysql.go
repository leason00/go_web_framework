package mymysql

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import mysql driver.
	"github.com/astaxie/beego/orm"
)

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql::url"))//密码为空格式
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}