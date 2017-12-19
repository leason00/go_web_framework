package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"fmt"
)

type User struct {
	Id          int
	Name        string
	Password	string
}

func ReadUser(username string)(password string, error error)  {
	o := orm.NewOrm()
	user := User{Name: username}
	err := o.Read(&user, "Name")
	if err == nil {
		return user.Password, nil
	}
	return "",  errors.New("查询不到！")
}


func ReadAllUser()([]User)  {
	o := orm.NewOrm()
	var users []User
	qs := o.QueryTable("user")
	cnt, err :=qs.Count()
	fmt.Println(cnt, err)
	res, err := qs.All(&users)
	fmt.Println(res)
	return users

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}