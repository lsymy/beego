package routers

import (
	"bgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.UserController{}, "post:Login")

}
