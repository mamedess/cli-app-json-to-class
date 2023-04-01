package langs

import (
	"encoding/json"
	"strconv"
)

var Str, ObjectName = "", ""

func decodeJSON(data []byte) map[string]interface{} {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		panic(err)
	}

	return jsonData
}

func Marshal(jsonData map[string]interface{}) {
	panic("unimplemented")
}

func isUniqueIn(slice []Prop, name string, father string) bool {
	for i := 0; i < len(slice); i++ {
		if strEquals(slice[i].name, name) && father != "" && strEquals(father, slice[i].father) {
			return false
		}
	}

	return true
}

func strEquals(str1 string, str2 string) bool {
	return str1 == str2
}

func GetType(v string) string {
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
