package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["titleString"] = "Home - My Golang Blog"
	c.Data["Islogin"] = checkAccount(c.Ctx)
	var err error
	c.Data["Topics"], err = models.GetAllTopic(true)
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "index.html"

}
