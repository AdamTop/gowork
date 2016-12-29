package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/replay", &controllers.ReplayController{})
	//beego自动路由
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.ReplayController{}) //留言自动路由
}
