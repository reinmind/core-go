package jwt

import (
	"time"

	"fmt"

	"github.com/golang-jwt/jwt"
	jwtgo "github.com/golang-jwt/jwt/v4"
)

func Decode(tokenString string) {
	// 整个匿名函数交给at去执行
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &Planet{},
			// 一个用于返回密码的函数
			// ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error)
			func(token *jwt.Token) (interface{}, error) {
				return []byte(nil), nil
			})

		if claims, ok := token.Claims.(*Planet); ok && token.Valid {
			fmt.Printf("method:%v\nname: %v\ndensity:%v\nradius:%v\nexpireAt:%v\n", token.Method.Alg(), claims.Name, claims.Density, claims.Radius, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Println(err)
		}
	})
}
func at(t time.Time, f func()) {
	// 构造一个返回t函数
	jwtgo.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwtgo.TimeFunc = time.Now
}
