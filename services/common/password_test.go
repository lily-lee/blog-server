package common

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncodePassword(t *testing.T) {
	Convey("Password", t, func() {
		salt := "aaa"
		password := "this is password"
		encodedStr := EncodePassword(salt, password)
		So(CheckPassword(salt, password, encodedStr), ShouldBeTrue)
		So(CheckPassword("bbb", password, encodedStr), ShouldBeFalse)
	})
}
