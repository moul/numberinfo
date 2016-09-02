package numberinfo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt64(t *testing.T) {
	Convey("Testing Int64()", t, func() {
		So(Int64(4242).int64, ShouldEqual, 4242)
		So(Int64(-4242).int64, ShouldEqual, -4242)
	})
}

func TestFloat64(t *testing.T) {
	Convey("Testing Float64()", t, func() {
		So(Float64(4242.0).float64, ShouldEqual, 4242.0)
		So(Float64(-4242.0).float64, ShouldEqual, -4242.0)
	})
}

func TestNumber_Int64(t *testing.T) {
	Convey("Testing Number.Int64()", t, func() {
		So(Int64(4242).Int64(), ShouldEqual, 4242)
		So(Int64(-4242).Int64(), ShouldEqual, -4242)
	})
}

func TestNumber_Float64(t *testing.T) {
	Convey("Testing Number.Float64()", t, func() {
		So(Int64(4242).Float64(), ShouldEqual, 4242)
		So(Int64(-4242).Float64(), ShouldEqual, -4242)
	})
}

func TestNumber_Int(t *testing.T) {
	Convey("Testing Number.Int()", t, func() {
		So(Int64(4242).Int(), ShouldEqual, 4242)
		So(Int64(-4242).Int(), ShouldEqual, -4242)
	})
}

func TestNumber_String(t *testing.T) {
	Convey("Testing Number.String()", t, func() {
		So(Int64(4242).String(), ShouldEqual, "4242")
		So(Int64(-4242).String(), ShouldEqual, "-4242")
	})
}
