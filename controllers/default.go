package controllers

import (
	"beego-demo/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	orm := models.GetOrmer()
	user := models.User{ID: 1}

	_ = orm.Read(&user)
	c.Data["Website"] = user.WebSite
	c.Data["Email"] = user.Email
	c.TplName = "index.tpl"
}
