package gokata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_boolparser(t *testing.T) {
	Convey("args unit test", t, func() {

		Convey("bp set_val with -l ", func() {
			bp := BooleanParser{}
			bp.setValue([]string{"-l"})
			So(bp.getValue(), ShouldEqual, true)
		})

		Convey("if -l not in input", func() {
			bp := BooleanParser{}
			bp.setValue([]string{""})
			So(bp.getValue(), ShouldEqual, false)
		})
	})
}

func Test_intparser(t *testing.T) {
	Convey("port unit test", t, func() {

		Convey("if -p in input with int", func() {
			it := IntParser{}
			it.setValue([]string{"-p", "6789"})
			So(it.getValue(), ShouldEqual, 6789)
		})

		Convey("if -p in input with word", func() {
			it := IntParser{}
			it.setValue([]string{"-p", "zxcxz"})
			So(it.getValue(), ShouldEqual, 0)
		})

		Convey("if -p in input with empty input", func() {
			it := IntParser{}
			it.setValue([]string{"-p", ""})
			So(it.getValue(), ShouldEqual, 0)
		})

		Convey("if with empty input", func() {
			it := IntParser{}
			it.setValue([]string{})
			So(it.getValue(), ShouldEqual, 0)
		})

	})
}

func Test_stringparser(t *testing.T) {
	Convey("port unit test", t, func() {

		Convey("if -d in input with path", func() {
			st := StringParser{}
			st.setValue([]string{"-d", "/usr/lib/local"})
			So(st.getValue(), ShouldEqual, "/usr/lib/local")
		})

		Convey("if -d in input with empty", func() {
			st := StringParser{}
			st.setValue([]string{"-d", ""})
			So(st.getValue(), ShouldEqual, "")
		})

		Convey("if input with empty", func() {
			st := StringParser{}
			st.setValue([]string{""})
			So(st.getValue(), ShouldEqual, "")
		})
	})

}

func Test_stringarrayparser(t *testing.T) {
	Convey("string_array_parser unit test", t, func() {

		Convey("if -g in input not with normal input ", func() {
			sap := StringArrayParser{}
			sap.setValue([]string{"-g", "asd", "asdsa", "zxcxz", "2321"})
			So(sap.getValue(), ShouldEqual, [50]string{"asd", "asdsa", "zxcxz", "2321"})
		})

		Convey("if -g in input not with empty input ", func() {
			sap := StringArrayParser{}
			sap.setValue([]string{"-g"})
			So(sap.getValue(), ShouldEqual, [50]string{""})
		})

		Convey("if  with empty input ", func() {
			sap := StringArrayParser{}
			sap.setValue([]string{})
			So(sap.getValue(), ShouldEqual, [50]string{""})
		})
	})
}

func Test_intarrayparser(t *testing.T) {
	Convey("stage2", t, func() {
		Convey("if -w in input not with normal input ", func() {
			iap := IntArrayParser{}
			iap.setValue([]string{"-w", "1", "2", "-3", "213"})
			So(iap.getValue(), ShouldEqual, [50]int{1, 2, -3, 213})
		})

		Convey("if -w in input not with string input ", func() {
			iap := IntArrayParser{}
			iap.setValue([]string{"-w", "1", "2", "-3", "213", "dfdasf"})
			So(iap.getValue(), ShouldEqual, [50]int{1, 2, -3, 213})
		})

		Convey("if -w in input not with empty input ", func() {
			iap := IntArrayParser{}
			iap.setValue([]string{"-w"})
			So(iap.getValue(), ShouldEqual, [50]int{0})
		})

		Convey("if with empty input ", func() {
			iap := IntArrayParser{}
			iap.setValue([]string{})
			So(iap.getValue(), ShouldEqual, [50]int{0})
		})
	})
}

func Test_schenaparse(t *testing.T) {
	Convey("test schema parse", t, func() {
		Convey("test boolparer in schema", func() {
			args := Args{marshaller: make(map[string]Parser)}
			args.schemaparse("l,")
			So(args.getBoolean("l"), ShouldEqual, false)
		})

		Convey("test intparer in schema", func() {
			args := Args{marshaller: make(map[string]Parser)}
			args.schemaparse("p,")
			So(args.getInt("p"), ShouldEqual, 0)
		})

		Convey("test stringparer in schema", func() {
			args := Args{marshaller: make(map[string]Parser)}
			args.schemaparse("d,")
			So(args.getString("d"), ShouldEqual, "")
		})

		Convey("test stringarrayparser in schema", func() {
			args := Args{marshaller: make(map[string]Parser)}
			args.schemaparse("g,")
			So(args.getStringarray("g"), ShouldEqual, [50]string{})
		})

		Convey("test intarrayparser in schema", func() {
			args := Args{marshaller: make(map[string]Parser)}
			args.schemaparse("w,")
			So(args.getIntarray("w"), ShouldEqual, [50]int{})
		})

	})
}

func Test_argparse(t *testing.T) {
	Convey("test arg parse", t, func() {
		//Convey("test boolparer in input args", func() {
		//	args := Args{marshaller: make(map[string]Parser)}
		//	args.schemaparse("l,")
		//	args.cmdparser([]string{"-l"})
		//	So(args.getBoolean("l"), ShouldEqual, true)
		//})

		// TODO to fix the common used case
		Convey("test string in input args", func() {
			args := Args{marshaller: make(map[string]Parser)}
			args.schemaparse("d, p")
			args.cmdparser([]string{"-d", "/usdsdaf"})
			So(args.getString("d"), ShouldEqual, "/usdsdaf")
		})

		//Convey("test int in input args", func() {
		//	args := Args{marshaller: make(map[string]Parser)}
		//	args.schemaparse("p,")
		//	args.cmdparser([]string{"-p", "5674"})
		//	So(args.getInt("l"), ShouldResemble, true)
		//})

		//Convey("test stringarray in input args", func() {
		//	args := Args{marshaller: make(map[string]Parser)}
		//	args.schemaparse("g,")
		//	args.cmdparser([]string{"-g", "sadsa", "/dfasf"})
		//	So(args.getStringarray("g"), ShouldResemble, true)
		//})
		//
		//Convey("test intarray in input args", func() {
		//	args := Args{marshaller: make(map[string]Parser)}
		//	args.schemaparse("w,")
		//	args.cmdparser([]string{"-w", "12313", "5212"})
		//	So(args.getIntarray("w"), ShouldResemble, true)
		//})
		//
		//
		//Convey("test intarray in input args", func() {
		//	args := Args{marshaller: make(map[string]Parser)}
		//	args.schemaparse("w,")
		//	args.cmdparser([]string{"-w", "12313", "5212"})
		//	So(args.getIntarray("w"), ShouldResemble, true)
		//})

	})
}
