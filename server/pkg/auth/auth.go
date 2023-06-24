package auth

import (
	"time"

	"github.com/cza14h/chat-nino-work/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var signer = []byte("chat.nino.work")

type AuthClaims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(userName string, userId uint) (string, error) {
	now := time.Now()
	expires := now.Add(config.JwtCookieExpiry)
	issuer := "chat.nino.work"

	claims := AuthClaims{
		Username: userName,
		UserID:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			Issuer:    issuer,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signer)
}

func ParseToken(c *gin.Context) (*AuthClaims, error) {
	token, err := c.Cookie(config.JwtTokenHeader)

	if err != nil {
		return nil, err
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return signer, nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AuthClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}
