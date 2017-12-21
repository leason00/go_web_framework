package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
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


func ReadAllUser(limit int, offset int)([]User,int64)  {
	o := orm.NewOrm()
	var users []User
	qs := o.QueryTable("user")
	total, _ := qs.Count()
	qs.Limit(limit, offset).All(&users)
	return users, total

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}