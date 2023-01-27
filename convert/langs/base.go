package langs

import (
	"encoding/json"
	"strconv"
)

var Str, Name = "", ""

func decodeJSON(data []byte) map[string]interface{} {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		panic(err)
	}

	return jsonData
}

func isUnique(name string, slice []Prop) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i].name == name {
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
