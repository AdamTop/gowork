package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
)

type ReplayController struct {
	beego.Controller
}

/*提交留言*/
func (this *ReplayController) Add() {
	//this.Ctx.WriteString(this.GetString("tid"))
	tid := this.GetString("tid")
	if len(tid) == 0 {
		this.Redirect("/topic", 302)
		return
	}
	name := this.GetString("name")
	email := this.GetString("email")
	content := this.GetString("content")
	_, err := models.Replayadd(tid, name, email, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/views/"+tid, 302)
}
