package routers

import (
	"bgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Include(&controllers.UsersController{})

	beego.Router("/login", &controllers.UserController{}, "post:Login")

}
