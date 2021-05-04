package routers

import (
	"beego-demo/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/users", &controllers.UserController{})
	beego.Router("/api/v1/users/:id", &controllers.UserController{}, "delete:UserDelete")
}
