package authorized

import (
	"errors"
	"synergy/web-service-gin/common/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var TokenExpireDuration time.Duration
var issuer string
var aud string
var MySecret []byte

func init() {
	config, _ := config.LoadConfig()
	MySecret = []byte(config.Secret)
	issuer = config.Issuer
	aud = config.Audience
	TokenExpireDuration = time.Hour * time.Duration(config.TokenExpire)
}

// GenToken generates JWT
func GenToken(username, email string) (string, error) {
	// Create our own statement
	c := MyClaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,                                      // subject of token
			IssuedAt:  time.Now().Unix(),                          // create token
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // expire token
			Issuer:    issuer,                                     // owner token
			Audience:  aud,                                        // who receive token
		},
	}
	// Creates a signed object using the specified signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Use the specified secret signature and obtain the complete encoded string token
	return token.SignedString(MySecret)
}

// ParseToken parsing JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // Verification token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
