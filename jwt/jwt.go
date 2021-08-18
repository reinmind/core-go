package jwt

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var hmacSampleSecret []byte

func GenSimple() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"foo":  "bar",
			"time": time.Now().Unix(),
		})
	s, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "signing error: %v", err)
	}
	return s
}
func Main() {
	fmt.Println(GenSimple())
}
