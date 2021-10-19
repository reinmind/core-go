package jwt

import (
	"fmt"
	"os"

	jwtgo "github.com/golang-jwt/jwt/v4"
)

var hmacSampleSecret []byte

func GenSimple() string {
	customClaims := Planet{
		2.5,
		"moon",
		0.25,
		[]string{"galaxy","solaris"},
		jwtgo.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, customClaims)
	s, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "signing error: %v", err)
	}
	return s
}
func Main() {
	fmt.Println(GenSimple())
}
