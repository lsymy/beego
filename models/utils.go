package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SECRET_KEY              string = "jwt-lg-test"
	DEFAULT_EXPIRED_SECONDS int    = 600
)

type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(loginInfo *LoginRequest, expiredSeconds int, userId int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRED_SECONDS
	}

	key := []byte(SECRET_KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	logs.Info("token expired at", time.Unix(expireAt,0))

	user := *loginInfo

	cl := jwt.StandardClaims{
		Issuer:    user.Username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expireAt,
	}

	claims := MyClaims{
		 userId,
		cl,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)

	if err != nil {
		errors.New("fail generateToken")
	}

	return tokenStr, nil
}
