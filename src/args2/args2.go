package args2

import (
	"strings"
)

type Args struct {
	parser     map[string]Parser
	argFound   map[string]bool
	currentarg [10]string
	pf         *ParserFactory
}

func (a *Args) schemaParse(input string) {
	schemas := strings.Split(input, ",")
	a.parser = make(map[string]Parser)
	a.pf = &ParserFactory{}
	for i, v := range schemas {
		getParser, option := a.pf.genParser(v)
		a.parser[option] = getParser
		a.currentarg[i] = option
	}
}

func (a *Args) parseOption(input [][]string) {
	for _, v := range input {
		if len(v) < 1 {
			continue
		} else {
			option := v[0]
			pureOption := strings.Split(option, "-")[1]
			a.parser[pureOption].setValue(v)
		}
	}
}

func (a *Args) getBoolean(param string) bool {
	return a.parser[param].getValue().(bool)
}

func (a *Args) getString(param string) string {
	return a.parser[param].getValue().(string)
}

func (a *Args) getInt(param string) int {
	return a.parser[param].getValue().(int)
}

func (a *Args) getIntArray(param string) [10]int {
	return a.parser[param].getValue().([10]int)
}

func (a *Args) getStringArray(param string) [10]string {
	return a.parser[param].getValue().([10]string)
}
