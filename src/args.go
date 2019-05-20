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
			a.marshaller[v] = &StringParser{}
		} else if v == "p" {
			a.marshaller[v] = &IntParser{}
		} else if v == "l" {
			a.marshaller[v] = &BooleanParser{}
		} else if v == "g" {
			a.marshaller[v] = &StringArrayParser{}
		} else if v == "w" {
			a.marshaller[v] = &IntArrayParser{}
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
	slow := 0
	fast := slow + 1
	for fast < 10 {
		if slow != 0 && (index[slow] == 0 || index[fast] == 0) {
			break
		} else {
			options := strings.Split(input[index[slow]], "-")
			if len(options) == 0 {
				break
			}
			option := options[1]
			if index[fast] == 0 {
				index[fast] = len(input)
			}
			value := input[index[slow]:index[fast]]
			a.marshaller[option].setValue(value)
			slow += 1
			fast += 1
		}
	}
}

func (a *Args) getBoolean(input string) bool {
	return a.marshaller[input].getValue().(bool)
}

func (a *Args) getInt(input string) int {
	return a.marshaller[input].getValue().(int)
}

func (a *Args) getString(input string) string {
	return a.marshaller[input].getValue().(string)
}

func (a *Args) getStringarray(input string) [50]string {
	return a.marshaller[input].getValue().([50]string)
}

func (a *Args) getIntarray(input string) [50]int {
	return a.marshaller[input].getValue().([50]int)
}
