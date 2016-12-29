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
	if len(name) == 0 || len(email) == 0 || len(content) == 0 {
		this.Redirect("/topic/views/"+tid, 302)
		return
	}
	_, err := models.Replayadd(tid, name, email, content)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/topic/views/"+tid, 302)
}

/*删除留言*/
func (this *ReplayController) Delete() {
	var err error
	nrid := this.GetString("rid")
	ntid := this.GetString("tid")
	if len(ntid) <= 0 {
		this.Redirect("/topic/views/"+ntid, 302)
		return
	}
	_, err = models.DeleteRepaly(nrid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/views/"+ntid, 302)
}
