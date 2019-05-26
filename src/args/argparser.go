package args

import "strconv"

type Parser interface {
	getValue() interface{}
	setValue(input []string)
}

type BooleanParser struct {
	val bool
}

func (bp *BooleanParser) getValue() interface{} {
	return bp.val
}

func (bp *BooleanParser) setValue(input []string) {
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

func (sp *StringParser) getValue() interface{} {
	return sp.val
}

func (sp *StringParser) setValue(input []string) {
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

func (ip *IntParser) getValue() interface{} {
	return ip.val
}

func (ip *IntParser) setValue(input []string) {
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

func (sap *StringArrayParser) getValue() interface{} {
	return sap.val
}

func (sap *StringArrayParser) setValue(input []string) {
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

func (iap *IntArrayParser) getValue() interface{} {
	return iap.val
}

func (iap *IntArrayParser) setValue(input []string) {
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

type ParseFactory struct{}

func (pf ParseFactory) createParser(i string) Parser {
	switch i {
	case "d":
		return &StringParser{}
	case "l":
		return &BooleanParser{}
	case "p":
		return &IntParser{}
	case "w":
		return &IntArrayParser{}
	case "g":
		return &StringArrayParser{}
	}
	return nil
}
