package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"apiproject/utils"
	"apiproject/msg"
)

// Operations about object
type UploadController struct {
	beego.Controller
}

func (this *UploadController) Upload() {
	method := this.Ctx.Request.Method
	if method == "GET"{
		this.Ctx.ResponseWriter.Header().Set("content-type", "text/html; charset=utf-8")
		this.Ctx.WriteString("<form id=\"form\" method=\"POST\" enctype=\"multipart/form-data\"><input id=\"myfile\" name=\"myfile\" type=\"file\" /><input type=\"submit\" value=\"保存\"  /></form>")
	}else{
		f, _, _ := this.GetFile("myfile")                  //获取上传的文件
		FileName := utils.GetUuid() + ".png"
		BasePath := beego.AppConfig.String("upload::url")
		path := BasePath + FileName							//文件目录
		f.Close()                                          		//关闭上传的文件，不然的话会出现临时文件不能清除的情况
		if utils.PathExists(BasePath){
			err := this.SaveToFile("myfile", path)
			fmt.Println(err)
			this.Data["json"] = msg.SuccessRes("上传成功！",  map[string]interface{}{"url":beego.AppConfig.String("upload::ResUrl") + FileName})
			this.ServeJSON()
		}
		this.Data["json"] = msg.ErrPath
		this.ServeJSON()

	}
}
