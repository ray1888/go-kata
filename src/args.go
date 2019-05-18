package gokata

import (
	"strings"
)

type Args struct {
	marshaller map[string]Parser
	args_found map[string]bool
	currentArg []string
}

func (a *Args) schemaparse(input string) {
	schema := strings.Split(input, ",")
	for _, v := range schema {
		if v == "d" {
			a.marshaller[v] = StringParser{}
		} else if v == "p" {
			a.marshaller[v] = IntParser{}
		} else if v == "l" {
			a.marshaller[v] = BooleanParser{}
		} else if v == "g" {
			a.marshaller[v] = StringArrayParser{}
		} else if v == "w" {
			a.marshaller[v] = IntArrayParser{}
		}
	}
}

func (a *Args) cmdparser(input []string) {
	var index [10]int
	var lastindex int
	for i, v := range input {
		if strings.HasPrefix(v, "-") {
			index[lastindex] = i
			lastindex += 1
		}
	}
	fast := 1
	slow := 0
	for fast < 10 {
		if slow != 0 && (index[slow] == 0 || index[fast] == 0) {
			break
		} else {
			option := strings.Split(input[slow], "-")
			value := input[index[slow]:index[fast]]
			a.marshaller[option[1]].set_value(value)
			slow += 1
			fast += 1
		}
	}
}

// TODO refactor to a factory mode

func (a *Args) get_boolean(input string) bool {
	return a.marshaller[input].get_value().(bool)
}

func (a *Args) get_int(input string) int {
	return a.marshaller[input].get_value().(int)
}

func (a *Args) get_string(input string) string {
	return a.marshaller[input].get_value().(string)
}

func (a *Args) get_stringarray(input string) [50]string {
	return a.marshaller[input].get_value().([50]string)
}

func (a *Args) get_intarray(input string) [50]int {
	return a.marshaller[input].get_value().([50]int)
}
