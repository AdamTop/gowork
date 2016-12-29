package controllers

import (
	"beeblog/models"
	"fmt"
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
	id := this.GetString("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	var err error
	if len(id) > 0 {
		_, err = models.UpdataTopic(id, title, content)
		if err != nil {
			beego.Error(err)
		} else {
			this.Redirect("/topic/views/"+id, 302)
		}
	} else {
		err = models.AddTopic(title, content)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/topic", 302)
	}

}

//获取文章显示id 数据
func (this *TopicController) Views() {
	var err error
	this.Data["titleString"] = "Article View- My Golang Blog"
	// if !checkAccount(this.Ctx) {
	// 	this.Redirect("/login", 302)
	// 	return
	// }
	this.Data["Islogin"] = checkAccount(this.Ctx)
	//this.Ctx.WriteString(this.Ctx.Input.Param("0"))
	//获取留言
	this.Data["Replays"], err = models.GetReplays(this.Ctx.Input.Param("0"))
	this.Data["Topics"], err = models.GetTopic(this.Ctx.Input.Param("0"))

	if err != nil {
		beego.Error(err)
		this.Redirect("/topic", 302)
		return
	}
	this.TplName = "topic_view.html"
}

//修改文章
func (this *TopicController) Edit() {
	var err error
	this.Data["titleString"] = "Article Edit- My Golang Blog"
	tid := this.Input().Get("tid")
	if len(tid) == 0 {
		this.Redirect("/topic", 302)
		return
	}
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.Data["Topics"], err = models.EditTopics(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/topic", 302)
		return
	}
	this.TplName = "topic_edit.html"
}

//删除文章
func (this *TopicController) Delete() {
	var err error
	sid := this.Input().Get("tid")
	if len(sid) == 0 {
		this.Redirect("/topic", 302)
		return
	}
	_, err = models.ArticleDel(sid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/topic", 302)
		return
	} else {
		fmt.Println("删除成功！")
		this.Redirect("/topic", 302)
		return
	}
}
