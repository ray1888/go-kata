package args2

import "strconv"

type Parser interface {
	setValue(input []string)
	getValue() interface{}
}

type BooleanParser struct {
	val bool
}

func (bp *BooleanParser) setValue(option []string) {
	if len(option) != 1 || option[0] == "" {
		bp.val = false
	} else {
		bp.val = true
	}
}

func (bp *BooleanParser) getValue() interface{} {
	return bp.val
}

type IntParser struct {
	val int
}

func (ip *IntParser) setValue(option []string) {
	if len(option) < 2 {
		return
	} else {
		number, err := strconv.Atoi(option[1])
		if err != nil {
			return
		}
		ip.val = number
	}
}

func (ip *IntParser) getValue() interface{} {
	return ip.val
}

type StringParser struct {
	val string
}

func (sp *StringParser) setValue(option []string) {
	if len(option) != 2 || option[1] == "" {
		return
	} else {
		sp.val = option[1]
	}
}

func (sp *StringParser) getValue() interface{} {
	return sp.val
}

type IntArrayParser struct {
	val [10]int
}

func (iap *IntArrayParser) setValue(option []string) {
	if len(option) <= 1 {
		return
	} else {
		option = option[1:]
		for i, v := range option {
			number, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			iap.val[i] = number
		}
	}
}

func (iap *IntArrayParser) getValue() interface{} {
	return iap.val
}

type StringArrayParser struct {
	val [10]string
}

func (sap *StringArrayParser) setValue(option []string) {
	if len(option) <= 1 {
		return
	} else {
		option = option[1:]
		for i, v := range option {
			sap.val[i] = v
		}
	}
}

func (sap *StringArrayParser) getValue() interface{} {
	return sap.val
}
