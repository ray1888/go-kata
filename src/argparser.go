package gokata

import "strconv"

type Parser interface {
	get_value() interface{}
	set_value(input []string)
}

type BooleanParser struct {
	val bool
}

func (bp BooleanParser) get_value() interface{} {
	return bp.val
}

func (bp BooleanParser) set_value(input []string) {
	// parse Bool parser
	if len(input) > 1 {
		bp.val = false
	}
	if input[0] == "-l" {
		bp.val = true
	} else {
		bp.val = false
	}
}

type StringParser struct {
	val string
}

func (sp StringParser) get_value() interface{} {
	return sp.val
}

func (sp StringParser) set_value(input []string) {
	if len(input) != 2 {
		return
	}
	if input[0] == "-d" {
		sp.val = input[1]
	}
}

type IntParser struct {
	val int
}

func (ip IntParser) get_value() interface{} {
	return ip.val
}

func (ip IntParser) set_value(input []string) {
	if len(input) != 2 {
		return
	}
	if input[0] == "-p" {
		val, err := strconv.Atoi(input[1])
		if err == nil {
			ip.val = val
		}
	}
}

type StringArrayParser struct {
	val [50]string
}

func (sap StringArrayParser) get_value() interface{} {
	return sap.val
}

func (sap StringArrayParser) set_value(input []string) {
	if len(input) == 0 {
		return
	}
	if input[0] == "-g" {
		input := input[1:]
		for i, v := range input {
			sap.val[i] = v
		}
	}
}

type IntArrayParser struct {
	val [50]int
}

func (iap IntArrayParser) get_value() interface{} {
	return iap.val
}

func (iap IntArrayParser) set_value(input []string) {
	if len(input) == 0 {
		return
	}
	if input[0] == "-w" {
		input := input[1:]
		for i, v := range input {
			val, err := strconv.Atoi(v)
			if err == nil {
				iap.val[i] = val
			}
		}
	}
}
