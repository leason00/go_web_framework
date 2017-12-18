package models

import (

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int
	Name        string
	Password	string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}