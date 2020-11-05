package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	err := orm.RegisterDriver("mysql", orm.DRMySQL)

	if err != nil {
		logs.Error(err.Error())
	}

	orm.RegisterModel(new(Users))

	dbuser := beego.AppConfig.String("username")
	dbpassword := beego.AppConfig.String("password")
	dbport := beego.AppConfig.String("port")
	dbhost := beego.AppConfig.String("host")
	dbname := beego.AppConfig.String("db")

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbuser,dbpassword,dbhost,dbport,dbname)
	fmt.Println(dburl)
	if err := orm.RegisterDataBase("default", "mysql", dburl); err != nil {
		logs.Error(err.Error())
	}

}
