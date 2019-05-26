package args2

import "regexp"

var intPattern *regexp.Regexp = regexp.MustCompile("[@]{1}")
var intArrayPattern *regexp.Regexp = regexp.MustCompile("[@][@]")
var stringPattern *regexp.Regexp = regexp.MustCompile("[!]{1}")
var stringArrayPattern *regexp.Regexp = regexp.MustCompile("[!][!]")
var booleanPattern *regexp.Regexp = regexp.MustCompile("#")
