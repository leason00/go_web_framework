package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//"apiproject/models"
)

// Operations about object
type TestController struct {
	beego.Controller
}

//var o = orm.NewOrm()


func (u *TestController) TestFunc() {
	//o.Using("default")
	//user := new(models.User)
	//user.Name = "slene"
	//
	//fmt.Println(o.Insert(user))

	//user.Name = "Your"
	//fmt.Println(o.Update(user))
	//fmt.Println(o.Delete(user))
	fmt.Println(888)
	u.Ctx.WriteString("jsoninfo is empty")
}

