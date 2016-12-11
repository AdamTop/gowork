package controllers

import (
	"beeblog/models"
	_ "fmt"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	action := this.Input().Get("action")
	switch action {
	case "del":
	case "edit":
	}
	var err error
	this.Data["Topics"], err = models.GetAllTopic(false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topic"] = true
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.Data["titleString"] = "Article - My Golang Blog"
	this.TplName = "topic.html"
}

//自动化路由方法测试
func (this *TopicController) Add() {
	//this.Ctx.WriteString("add router")
	this.Data["titleString"] = "Article Add- My Golang Blog"
	this.TplName = "topic_add.html"
}

//获取全部post过来的数据方法
func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)

}

//获取文章显示id 数据
func (this *TopicController) Views() {
	this.Data["titleString"] = "Article View- My Golang Blog"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.Data["Islogin"] = checkAccount(this.Ctx)
	//fmt.Println(this.Ctx.Input.Bind(&id, "id"))
	var id string
	var vid string
	var err error
	vid, err = this.Ctx.Input.Bind(&id, "id")
	if err != nil {
		this.Error(err)
		this.Redirect("/topic", 302)
		return
	}
	/*
		this.Data["Topic"] = models.GetTopic(this.Ctx.Input.Params("0"))
		if err != nil {
			this.Error(err)
			this.Redirect("/topic", 302)
			return
		}*/
	this.TplName = "topic_view.html"
}
