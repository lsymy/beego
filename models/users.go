package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"net/http"
)

type Users struct {
	Id       int64
	Username string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
	Password string `orm:"size(128)"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	UserId   int    `json:"id"`
	Token    string `json:"token"`
}

//func init()  {
//	orm.RegisterModel(new(Users))
//}

func Login(lr *LoginRequest) (*LoginResponse, int, error) {
	username := lr.Username
	password := lr.Password

	if len(username) == 0 || len(password) == 0 {
		return nil, http.StatusBadRequest, errors.New("username or password null")
	}

	o := orm.NewOrm()

	user := &Users{Username: username}
	err := o.Read(user, "username")
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	tokenString , err := GenerateToken(lr, 0, (int(user.Id)))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &LoginResponse{
		Username: username,
		UserId: int(user.Id),
		Token: tokenString,
	}, http.StatusOK, err
}
