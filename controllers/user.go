package controllers

import (
	"beego-demo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

// UserInfo request user create/update
type UserInfo struct {
	ID      int64  `json:"id,omitempty"`
	User    string `json:"user,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	WebSite string `json:"web_site,omitempty"`
	Address string `json:"address,omitempty"`
}

func (c *UserController) Get() {

	orm := models.GetOrmer()
	userItems := []models.User{}
	_, err := orm.QueryTable("user").All(&userItems)
	if err != nil {
		logs.Error("get error: %v", err.Error())
		http.Error(c.Ctx.ResponseWriter, err.Error(), 500)
	}
	c.Data["json"] = userItems
	c.ServeJSON()
}

func (c *UserController) Post() {
	request := UserInfo{}

	// TODO: add
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err == nil {
	// 	logs.Error("request parse err: %v", err.Error())
	// }

	err := json.Unmarshal(c.Ctx.Input.CopyBody(1<<32), &request)
	if err != nil {
		logs.Error("Invalid json request: " + err.Error())
		return
	}
	logs.Debug("request: %+v", request)

	orm := models.GetOrmer()
	user := models.User{
		Name:    request.User,
		User:    request.User,
		Email:   request.Email,
		Address: request.Address,
		WebSite: request.WebSite,
	}

	id, err := orm.Insert(&user)
	if err == nil {
		logs.Debug("create user id: %v", id)
	}

	c.Data["json"] = map[string]interface{}{"id": id}
	c.ServeJSON()
}

func (c *UserController) Put() {
	request := UserInfo{}
	err := json.Unmarshal(c.Ctx.Input.CopyBody(1<<32), &request)
	if err != nil {
		logs.Error("Invalid json request: " + err.Error())
		return
	}
	logs.Debug("user update request: %+v", request)

	if request.ID == 0 {
		http.Error(c.Ctx.ResponseWriter, "请求不合法", 400)
	}

	orm := models.GetOrmer()
	user := models.User{ID: int(request.ID)}
	err = orm.Read(&user)
	if err == nil {
		user.Name = request.Name
		user.User = request.User
		user.Email = request.Email
		user.Address = request.Address
		user.WebSite = request.WebSite
	} else {
		logs.Error("update error: %v", err.Error())
		http.Error(c.Ctx.ResponseWriter, err.Error(), 500)
	}

	if num, err := orm.Update(&user); err == nil {
		fmt.Println(num)
	}

	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) UserDelete() {
	value := c.Ctx.Input.Param(":id")
	userID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		logs.Error("when delete user, get user id occur error: %v", err.Error())
		http.Error(c.Ctx.ResponseWriter, err.Error(), 500)
	}
	logs.Debug("user update request, userID: %v", userID)

	if userID == 0 {
		http.Error(c.Ctx.ResponseWriter, "请求不合法", 400)
	}

	orm := models.GetOrmer()
	user := models.User{ID: int(userID)}
	if num, err := orm.Delete(&user); err == nil {
		fmt.Println(num)
	}

	c.Data["json"] = user
	c.ServeJSON()
}
