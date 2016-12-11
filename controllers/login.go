package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//return
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.Data["titleString"] = "Login - My Golang Blog"
	this.TplName = "login.html"
}
func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//true
	uname := this.Input().Get("uname")
	upwd := this.Input().Get("pwd")
	autologin := this.Input().Get("autologin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("password") == upwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1 //大的int型
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", upwd, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("password") == pwd
}
