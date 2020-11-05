package controllers

import (
	"bgo/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Login()  {
	lr := new(models.LoginRequest)

	u.ParseForm(lr)

	loginResponse, statusCode, err := models.Login(lr)
	if err != nil {
		u.respond(statusCode, err.Error())
		return
	}

	u.Ctx.Output.Header("Authorization", loginResponse.Token)
	u.respond(statusCode, "", loginResponse)
}

func (u *UserController) respond(code int, message string, data ...interface{}) {
	u.Ctx.Output.SetStatus(code)

	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	u.Data["json"] = struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Code:    code,
		Message: message,
		Data:    d,
	}
	u.ServeJSON()
}