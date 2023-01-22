package langs

import (
	"main/util"
	"strconv"
)

var Str, Name = "", ""

func GetPropNames(prompt string) []Prop {
	//get a slice with the names of the properties
	pattern := util.GetPattern(`"([^"]*)"`)
	matches := pattern.FindAllStringSubmatch(prompt, -1)

	um := []Prop{}
	lm := len(matches)
	p := Prop{}

	for i := 0; i < lm; i++ {
		if p.value != "" {
			p = Prop{}
		}

		v := matches[i][1]

		if i%2 == 0 {
			p.name = v
		} else {
			p.value = v
		}

		if p.value != "" && isUnique(p, um) {
			um = append(um, p)
		}
	}

	return um
}

func isUnique(prop Prop, slice []Prop) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i].name == prop.name {
			return false
		}
	}

	return true
}

func GetValueType(v string) string {
	_, err := strconv.ParseInt(v, 10, 64)

	if err == nil {
		return "int64"
	}
	_, err = strconv.ParseBool(v)

	if err == nil {
		return "bool"
	}
	_, err = strconv.ParseFloat(v, 64)

	if err == nil {
		return "float"
	}

	return "string"
}

// func trimStr(str string, cutstr string) string {
// 	return strings.Trim(str, cutstr)
// }
