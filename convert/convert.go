package convert

import (
	"bufio"
	"regexp"
)

// variable with all the properties of the json
var jsonProp = []string{}

func getPropNames(prompt string, r *bufio.Reader) {
	//get a slice with the whole property and it's value in the format of a string
	pattern := getPattern("ck$")
	m := pattern.Split(prompt, -1)

	//get a slice with the names of the properties
	pattern = getPattern("ck$")
	n := pattern.Split(prompt, -1)

	for i := 0; i < len(m); i++ {
		jsonProp = append(jsonProp, n[i])
	}
}

func JsonToGolang(prompt string, r *bufio.Reader) {
	getPropNames(prompt, r)
	// sb := CreateBaseJson()
	for i := 0; i < len(jsonProp); i++ {
		// appendTo(sb)
	}
}

func JsonToCSharp(prompt string, r *bufio.Reader)

func getPattern(regex string) *regexp.Regexp {
	return regexp.MustCompile(regex)
}
