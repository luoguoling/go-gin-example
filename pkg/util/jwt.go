package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-gin-example/pkg/setting"
	"time"

	//"time"
)

var jwtSecret = []byte(setting.JwtSecret)
const TokenExpireDuration = time.Hour * 2

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username,password string) (string,error) {
	fmt.Println("产生token",username,password)
	fmt.Println(jwtSecret)
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: expireTime.Unix(),
			Issuer: "rolin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err := tokenClaims.SignedString(jwtSecret)
	return token,err
}
func ParseToken(token string) (*Claims,error)  {
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{},error) {
		return jwtSecret,nil
	})
	if tokenClaims != nil{
		if claims,ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err
}
