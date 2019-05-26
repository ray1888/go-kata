package args2

import (
	"regexp"
	"strings"
	//"regexp"
)

type Args struct {
	parser     map[string]Parser
	argFound   map[string]bool
	currentarg [10]string
}

func (a *Args) schemaParse(input string) {
	schemas := strings.Split(input, ",")
	a.parser = make(map[string]Parser)
	for i, v := range schemas {
		if match, _ := regexp.MatchString(booleanPattern, v); match {
			option := strings.Split(v, "#")
			o := option[0]
			a.parser[o] = &BooleanParser{}
			a.currentarg[i] = o
		} else if match, _ := regexp.MatchString(stringArrayPattern, v); match {
			option := strings.Split(v, "!")[0]
			a.parser[option] = &StringArrayParser{}
			a.currentarg[i] = option
		} else if match, _ := regexp.MatchString(stringPattern, v); match {
			option := strings.Split(v, "!")[0]
			a.parser[option] = &StringParser{}
			a.currentarg[i] = option
		} else if match, _ := regexp.MatchString(intArrayPattern, v); match {
			option := strings.Split(v, "@")[0]
			a.parser[option] = &IntArrayParser{}
			a.currentarg[i] = option
		} else if match, _ := regexp.MatchString(intPattern, v); match {
			option := strings.Split(v, "@")[0]
			a.parser[option] = &IntParser{}
			a.currentarg[i] = option
		}
	}
}

func (a *Args) parseOption(input [][]string) {
	for _, v := range input {
		if len(v) < 1 {
			continue
		} else {
			option := v[0]
			optionUndash := strings.Split(option, "-")[1]
			a.parser[optionUndash].setValue(v)
		}
	}
}

func (a *Args) findTypeExist() {

}

// TODO this underneath part can be restructure
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
