package jwt

import jwtgo "github.com/golang-jwt/jwt/v4"

type Planet struct {
	Radius  float64 `json:"radius"`
	Name    string  `json:"name"'`
	Density float32 `json:"density"`
	Position []string `json:"position"`
	jwtgo.StandardClaims
}
