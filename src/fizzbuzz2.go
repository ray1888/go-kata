package gokata

import "strconv"

// func fizzbuzz_count_input(number int) {
// 	times := number / 100
// 	length := number % 100

// 	results := [10000]string{}
// 	fizzbuzz100 := fizzbuzzto100()
// 	var start int = 0
// 	for start; start < times-1; start++ {
// 		results
// 	}
// }

func fizzbuzzto100() []string {
	results := make([]string, 100)
	var i int
	for i = 0; i < 100; i++ {
		results[i] = fizzbuzz2(i)
	}
	return results
}

func fizzbuzz2(number int) string {
	var result string = ""
	if number != 0 && number%3 == 0 && number%5 == 0 {
		result += "FizzBuzz"
		return result
	}
	var isFizz string = isfizz(number)
	if isFizz != "" {
		result += isFizz
	}
	var isBuzz string = isbuzz(number)
	if isbuzz(number) != "" && isFizz == "" {
		result += isBuzz
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}

func isfizz(number int) string {
	if number != 0 && (number%3 == 0 || number/10 == 3 || number%10 == 3) {
		return "Fizz"
	} else {
		return ""
	}
}

func isbuzz(number int) string {
	if number != 0 && (number%5 == 0 || number/10 == 5 || number%10 == 5) {
		return "Buzz"
	} else {
		return ""
	}
}
