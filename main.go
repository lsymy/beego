package main

import (
	_ "bgo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}

