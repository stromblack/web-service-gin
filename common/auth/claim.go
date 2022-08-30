package authorized

import (
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
