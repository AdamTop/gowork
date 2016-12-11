package controllers

import (
	"beeblog/models"
	//"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	action := this.Input().Get("action")
	switch action {
	case "add":
		//this.Ctx.WriteString(fmt.Sprint(this.Input()))
		//return
		catename := this.Input().Get("catename")
		if len(catename) == 0 {
			break
		}
		err := models.AddCategory(catename)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		cateId := this.Input().Get("id")
		if len(cateId) == 0 {
			break
		}
		err := models.DelCategory(cateId)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}

	this.Data["Category"] = true
	this.Data["titleString"] = "Category - My Golang Blog"
	var err error
	this.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.TplName = "category.html"
}

func (this *CategoryController) Post() {

}
