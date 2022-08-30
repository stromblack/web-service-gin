package authorized

import (
	"synergy/web-service-gin/models"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	*jwt.RegisteredClaims
	UserInfo models.User
}
