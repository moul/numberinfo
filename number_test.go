package numberinfo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt64(t *testing.T) {
	Convey("Testing Int64()", t, func() {
		number := Int64(4242)
		So(number.Int64(), ShouldEqual, 4242)
	})
}
