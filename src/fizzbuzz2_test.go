package gokata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {

	Convey("FizzBuzz testing unit", t, func() {
		Convey("test fizz", func() {
			Convey("should mod by 3 ", func() {
				So(fizzbuzz2(3), ShouldEqual, "Fizz")
			})

			Convey("should not mod by 3 ", func() {
				So(fizzbuzz2(2), ShouldEqual, "2")
			})

			Convey("should contain 3 ", func() {
				So(fizzbuzz2(13), ShouldEqual, "Fizz")
			})

			Convey("should contain 3 in second ", func() {
				So(fizzbuzz2(31), ShouldEqual, "Fizz")
			})

			Convey("input 0 should return 0", func() {
				So(fizzbuzz2(0), ShouldEqual, "0")
			})

		})

		Convey("test buzz", func() {
			Convey("should mod by 5 ", func() {
				So(fizzbuzz2(5), ShouldEqual, "Buzz")
			})

			Convey("should not mod by 5", func() {
				So(fizzbuzz2(2), ShouldEqual, "2")
			})

			Convey("should contain 5", func() {
				So(fizzbuzz2(25), ShouldEqual, "Buzz")
			})

			Convey("should contain 5 in second ", func() {
				So(fizzbuzz2(56), ShouldEqual, "Buzz")
			})
		})

		Convey("test fizzbuzz2", func() {
			Convey("should mod by 5 ,3 ", func() {
				So(fizzbuzz2(15), ShouldEqual, "FizzBuzz")
			})

			Convey("should not mod by 5 and 3", func() {
				So(fizzbuzz2(7), ShouldEqual, "7")
			})
		})

		// Convey("test range 100", func() {
		// 	Convey("test range 100 ", func() {
		// 		So(fizzbuzzto100(), ShouldEqual, "[0 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz Fizz 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 Fizz Fizz Buzz 26 Fizz 28 29 FizzBuzz Fizz Fizz Fizz Fizz Fizz Fizz Fizz Fizz Fizz Buzz 41 Fizz Fizz 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz Buzz Fizz Fizz Buzz Buzz Fizz Buzz Buzz FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz Fizz 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 Fizz Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz]")
		// 	})

		// })

	})

}
