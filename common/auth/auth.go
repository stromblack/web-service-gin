package authorized

import (
	"errors"
	"fmt"
	"synergy/web-service-gin/common/config"
	"synergy/web-service-gin/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var TokenExpireDuration time.Duration
var issuer string
var aud string
var MySecret []byte

const RefreshTokenExpire = time.Hour * 8

var MyRefreshSecret = []byte("synergysoftware2008")

func init() {
	config, _ := config.LoadConfig()
	MySecret = []byte(config.Secret)
	issuer = config.Issuer
	aud = config.Audience
	TokenExpireDuration = time.Hour * time.Duration(config.TokenExpire)
}

// GenToken generates JWT
func GenToken(userinfo models.User) (string, error) {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(TokenExpireDuration)
	iat := time.Now()
	// Add your claims
	token.Claims = &MyClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the unque value
			Issuer:    issuer,
			Audience:  []string{aud},
			Subject:   userinfo.Email,
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(iat),
		},
		userinfo,
	}
	// Sign the token with your secret key
	return token.SignedString(MySecret)
}

// ParseToken parsing JWT
func VerifyToken(tokenString string) (*MyClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
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

// GenToken generates JWT
func RefreshToken(userinfo models.User) (string, error) {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(RefreshTokenExpire)
	iat := time.Now()
	// Add your claims
	token.Claims = &MyClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the unque value
			Issuer:    issuer,
			Audience:  []string{aud},
			Subject:   userinfo.Email,
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(iat),
		},
		models.User{},
	}
	// Sign the token with your secret key
	return token.SignedString(MyRefreshSecret)
}
func VerifyRefreshToken(tokenString string) (*MyClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return MyRefreshSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // Verification token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
