package args2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Parser(t *testing.T) {
	Convey("test boolean parser", t, func() {

		Convey("bp set_val with -l ", func() {
			bp := &BooleanParser{}
			bp.setValue([]string{"-l"})
			So(bp.getValue(), ShouldEqual, true)
		})

		Convey("if -l not in input", func() {
			bp := BooleanParser{}
			bp.setValue([]string{""})
			So(bp.getValue(), ShouldEqual, false)
		})
	})

	Convey("test int parser", t, func() {

		Convey("ip set_val with val ", func() {
			ip := IntParser{}
			ip.setValue([]string{"-p", "8080"})
			So(ip.getValue(), ShouldEqual, 8080)
		})

		Convey("ip set_val with empty", func() {
			ip := IntParser{}
			ip.setValue([]string{""})
			So(ip.getValue(), ShouldEqual, 0)
		})
	})

	Convey("test string parser", t, func() {

		Convey("bs set_val with -d ", func() {
			sp := &StringParser{}
			sp.setValue([]string{"-d", "/test"})
			So(sp.getValue(), ShouldEqual, "/test")
		})

		Convey("if -l not in input", func() {
			sp := StringParser{}
			sp.setValue([]string{"-d"})
			So(sp.getValue(), ShouldEqual, "")
		})
	})

	Convey("test int array parser", t, func() {

		Convey("iap set_val with -g ", func() {
			iap := &IntArrayParser{}
			iap.setValue([]string{"-g", "5050", "6060"})
			So(iap.getValue(), ShouldEqual, [10]int{5050, 6060})
		})

		Convey("if -l not in input", func() {
			iap := &IntArrayParser{}
			iap.setValue([]string{"-g"})
			So(iap.getValue(), ShouldEqual, [10]int{})
		})
	})

	Convey("test string array parser", t, func() {

		Convey("sap set_val with -w", func() {
			sap := &StringArrayParser{}
			sap.setValue([]string{"-w", "/test", "/usr/bin"})
			So(sap.getValue(), ShouldEqual, [10]string{"/test", "/usr/bin"})
		})

		Convey("if -l not in input", func() {
			sap := &StringArrayParser{}
			sap.setValue([]string{"-w"})
			So(sap.getValue(), ShouldEqual, [10]string{})
		})
	})
}
