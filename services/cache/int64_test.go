package cache

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/lily-lee/blog-server/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInt64(t *testing.T) {
	config.InitConfig("../../.env")
	config.InitRedis()
	rc := config.RC
	Convey("TestGetInt64", t, func() {
		num := time.Now().Unix()
		key := "test_get_int64_not_exists"
		rc.Del(context.Background(), key)

		val, err := rc.Get(context.Background(), key).Int64()
		So(err, ShouldEqual, redis.Nil)
		So(val, ShouldEqual, 0)

		val, err = GetInt64(key, func() (int64, error) {
			return num, nil
		})
		So(err, ShouldBeNil)
		So(val, ShouldEqual, num)

		val, err = rc.Get(context.Background(), key).Int64()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, num)
	})
}
