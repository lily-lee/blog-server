package common

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandom(t *testing.T) {
	Convey("Random", t, func() {
		for i := 10; i <= 32; i++ {
			str := Random(i)
			So(len(str), ShouldEqual, i)
		}
	})
}
