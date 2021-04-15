package controllers

import (
	"beego-demo/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	orm := models.GetOrmer()

	value := c.Ctx.Input.Query("id")
	v, _ := strconv.ParseInt(value, 10, 64)

	user := models.User{ID: int(v)}

	err := orm.Read(&user)
	if err != nil {
		c.Data["Name"] = "404"
	} else {
		c.Data["Website"] = user.WebSite
		c.Data["Email"] = user.Email
		c.Data["Name"] = user.Name
	}
	c.TplName = "index.tpl"
}
