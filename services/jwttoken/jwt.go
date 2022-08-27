package jwttoken

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/request"
)

type UserClaims struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type Token struct {
	Signature string `json:"signature"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiresAt int64  `json:"expires_at"`
}

// GenToken generate token
func GenToken(user *models.User) (*Token, error) {
	claims := &UserClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Email,
			ExpiresAt: time.Now().Add(config.Conf.TokenDuration).Unix(),
			Id:        fmt.Sprintf("%d", user.ID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    user.Email,
			Subject:   "login",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signStr, err := t.SignedString([]byte(config.Conf.SignKey))
	if err != nil {
		return nil, err
	}

	token := &Token{
		Signature: signStr,
		IssuedAt:  claims.IssuedAt,
		ExpiresAt: claims.ExpiresAt,
	}

	return token, nil
}

func ParseToken(tokenStr string) (*UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, &request.BizErr{HttpCode: http.StatusUnauthorized, ErrMsg: "invalid token"}
	}

	return claims, nil
}

type UserInfo struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

func GetUser(c *gin.Context) (*UserInfo, error) {
	claims, ok := c.MustGet("user_claims").(*UserClaims)
	if !ok || claims == nil {
		return nil, &request.BizErr{HttpCode: http.StatusUnauthorized}
	}

	return &UserInfo{ID: claims.ID, Email: claims.Email}, nil
}
