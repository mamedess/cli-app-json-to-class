package langs

import (
	"fmt"
	"main/util"
)

// variable with all the properties of the json
var jsonProps = []string{}
var Str = ""

func GetPropNames(prompt string) []string {

	fmt.Println(prompt)
	//get a slice with the names of the properties
	pattern := util.GetPattern(`"([^"]*)":`)
	matches := pattern.FindAllStringSubmatch(prompt, -1)

	for _, match := range matches {
		jsonProps = append(jsonProps, match[1])
	}

	return jsonProps
}
