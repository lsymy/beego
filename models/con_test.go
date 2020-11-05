package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"testing"
)

func TestCon(T *testing.T)  {
	fmt.Println(beego.AppConfig.String("username"))
}
