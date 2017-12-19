package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"apiproject/models"
	_ "apiproject/msg"
	"apiproject/msg"
	"fmt"
)

// Operations about object
type UserController struct {
	beego.Controller
}


type LoginData struct {
	Username    string	`json:"username"`
	Password	string	`json:"password"`
}

func (u *UserController) LoginFunc() {
	var ob LoginData//这是一个model，struct类型
	body := u.Ctx.Input.RequestBody//这是获取到的json二进制数据
	json.Unmarshal(body, &ob)//解析二进制json，把结果放进ob中
	password, err := models.ReadUser(ob.Username)
	if err != nil {
		u.Data["json"] = msg.ErrorRes("账号不存在", 1)
		u.ServeJSON()
	}
	if password != ob.Password{
		u.Data["json"] = msg.ErrorRes("密码错误！", 1)
		u.ServeJSON()
	}
	u.Data["json"] = msg.SuccessRes("登录成功！", map[string]interface{}{"token": "51664164165"})
	u.ServeJSON()
}

func (u *UserController) ListFunc() {

	res := models.ReadAllUser()
	data := make([]interface{}, 0)
	for _, value := range res {
		data = append(data, map[string]interface{}{"id": value.Id, "username": value.Name})
	}
	fmt.Println(data)
	u.Data["json"] = msg.ArrayRes("查询成功！", data)
	u.ServeJSON()
}