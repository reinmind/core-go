package jwt

import jwtgo "github.com/golang-jwt/jwt/v4"

type Planet struct {
	Radius  float32 `json:2.5`
	Name    string  `json:"earth"`
	Density float32 `json:0.3`
	jwtgo.StandardClaims
}
