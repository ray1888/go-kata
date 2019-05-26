package args2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Schema(t *testing.T) {
	Convey("test basic three types", t, func() {
		args := &Args{}
		args.schemaParse("l#,d!,p@")
		args.parseOption([][]string{
			[]string{"-l"},
			[]string{"-p", "5050"},
			[]string{"-d", "33333"},
		})

		So(args.getBoolean("l"), ShouldEqual, true)
		So(args.getInt("p"), ShouldEqual, 5050)
		So(args.getString("d"), ShouldEqual, "33333")
	})

	Convey("test additional two type", t, func() {
		args := &Args{}
		args.schemaParse("l#,w!!,g@@")
		args.parseOption([][]string{
			[]string{"-l"},
			[]string{"-g", "5950", "6030"},
			[]string{"-w", "/usr/lib", "/tmp/akb"},
		})
		So(args.getIntArray("g"), ShouldEqual, [10]int{5950, 6030})
		So(args.getStringArray("w"), ShouldEqual, [10]string{"/usr/lib", "/tmp/akb"})
	})

}
