package langs

import "main/util"

// variable with all the properties of the json
var jsonProp = []string{}
var Str = ""

func GetPropNames(prompt string) []string {
	//get a slice with the whole property and it's value in the format of a string
	pattern := util.GetPattern("ck$")
	m := pattern.Split(prompt, -1)

	//get a slice with the names of the properties
	pattern = util.GetPattern("ck$")
	n := pattern.Split(prompt, -1)

	for i := 0; i < len(m); i++ {
		jsonProp = append(jsonProp, n[i])
	}

	return jsonProp
}
