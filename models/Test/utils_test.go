package Test

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"

	//"time"
)

type MyClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func TestGenerateToken(T *testing.T)  {
	expiredSeconds := 600
	SECRET_KEY := "jwt"
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	key := []byte(SECRET_KEY)
	fmt.Println(key)
	fmt.Println(time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix())

	cl := jwt.StandardClaims{
		Issuer:    "admin1",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expireAt,
	}

	claims := MyClaims{
		1,
		cl,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)

	if err != nil {
		errors.New("fail generateToken")
	}

fmt.Println(tokenStr)
}

