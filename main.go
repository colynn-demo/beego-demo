package main

import (
	"beego-demo/models"
	_ "beego-demo/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	models.DBinit()
	beego.Run()
}
