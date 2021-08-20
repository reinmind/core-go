package jwt

import (
	"encoding/json"
	"fmt"
	jwtgo "github.com/golang-jwt/jwt/v4"
	"time"
)

func Decode(tokenString string) string {
	var s string
	// 整个匿名函数交给at去执行
	at(time.Unix(0, 0), func() {
		token, err := jwtgo.ParseWithClaims(tokenString, &Planet{},
			// 一个用于返回密码的函数
			// ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error)
			func(token *jwtgo.Token) (interface{}, error) {
				return []byte(nil), nil
			})

		if claims, ok := token.Claims.(*Planet); ok && token.Valid {
			marshal, err := json.Marshal(claims)
			if err == nil {
				s = fmt.Sprintf(
					"claims:%v\n", marshal)
			}
		} else {
			fmt.Println(err)
		}
	})
	return s
}
func at(t time.Time, f func()) {
	// 构造一个返回t函数
	jwtgo.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwtgo.TimeFunc = time.Now
}
