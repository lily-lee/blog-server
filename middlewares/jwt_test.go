package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
)

func TestJwtAuth(t *testing.T) {
	config.InitConfig("../.env")
	type testcase struct {
		JwtToken   string
		ExpectCode int
	}
	token, _ := jwttoken.GenToken(&models.User{
		ID:    1,
		Email: "www@gmail.com",
	})
	cases := []testcase{
		{ExpectCode: 401},
		{JwtToken: "xxx", ExpectCode: 401},
		{JwtToken: token.Signature, ExpectCode: 200},
	}
	Convey("JwtAuth", t, func() {
		router := gin.New()
		router.Use(JwtAuth())
		router.POST("/user")
		server := httptest.NewServer(router)
		for _, c := range cases {
			req, err := http.NewRequest(http.MethodPost, server.URL+"/user", nil)
			So(err, ShouldBeNil)

			req.Header.Add("Authorization", c.JwtToken)
			client := http.DefaultClient
			resp, err := client.Do(req)

			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, c.ExpectCode)
		}
	})
}
