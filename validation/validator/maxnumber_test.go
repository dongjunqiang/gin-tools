package validator

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxNumberValidator(t *testing.T) {

	threshold := 1000

	Convey("When creating a MaxNumber validator", t, func() {
		v := MaxNumber{threshold}
		Convey("give 999 paramater", func() {
			So(v.Validate("999"), ShouldBeTrue)
		})
		Convey("give 1000 paramater", func() {
			So(v.Validate("1000"), ShouldBeTrue)
		})
		Convey("give 1001 paramater", func() {
			So(v.Validate("1001"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeFalse)
		})
	})
	Convey("When creating a MaxNumber if not empty validator", t, func() {
		v := MaxNumberIfNotEmpty{threshold}
		Convey("give 999 paramater", func() {
			So(v.Validate("999"), ShouldBeTrue)
		})
		Convey("give 1000 paramater", func() {
			So(v.Validate("1000"), ShouldBeTrue)
		})
		Convey("give 1001 paramater", func() {
			So(v.Validate("1001"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})
}
