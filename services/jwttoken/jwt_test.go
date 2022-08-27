package jwttoken

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
)

func TestGenToken(t *testing.T) {
	config.InitConfig("../../.env")
	Convey("Jwt Token", t, func() {
		user := &models.User{
			ID:    1,
			Email: "www@gmail.com",
		}

		Convey("success", func() {
			token, err := GenToken(user)
			So(err, ShouldBeNil)
			So(token, ShouldNotBeNil)

			claims, err := ParseToken(token.Signature)
			So(err, ShouldBeNil)
			So(claims, ShouldNotBeNil)
			So(claims.ID, ShouldEqual, user.ID)
			So(claims.Email, ShouldEqual, user.Email)
		})

		Convey("signature invalid", func() {
			claims := &UserClaims{
				ID: user.ID, Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					IssuedAt:  time.Now().Unix(),
					ExpiresAt: 0,
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			s, err := token.SignedString([]byte("aaaa"))
			So(err, ShouldBeNil)

			c, err := ParseToken(s)
			So(c, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "signature is invalid")
		})
	})
}
