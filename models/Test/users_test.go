package Test

import (
	"bgo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testing"
)

func TestLogin(T *testing.T)  {
	username := "admin1"
	password := "1234567"
fmt.Println(beego.AppConfig.String("username"))
	o := orm.NewOrm()
	user := &models.Users{Username: username}
	err := o.Read(user, "username")
	if err != nil {
		fmt.Println("err")
	}

	lr := &models.LoginRequest{Username: username, Password: password}
	tokenStr, err := models.GenerateToken(lr, 600, 1)
	fmt.Println(tokenStr)
}
