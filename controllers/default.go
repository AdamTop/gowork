package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var err error
	c.Data["IsHome"] = true
	c.Data["titleString"] = "Home - My Golang Blog"
	c.Data["Islogin"] = checkAccount(c.Ctx)
	c.Data["Categorys"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	cid := c.Input().Get("cateid")
	if len(cid) > 0 {
		c.Data["Topics"], err = models.GetAllTopic(true, cid)
		if err != nil {
			beego.Error(err)
		}
	} else {
		c.Data["Topics"], err = models.GetAllTopic(true, " ")
		if err != nil {
			beego.Error(err)
		}
	}

	c.TplName = "index.html"

}
