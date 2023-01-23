package langs

import (
	"encoding/json"
)

var Str, Name = "", ""

func decodeJSON(data []byte) (string, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		return "", err
	}

	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

// func trimStr(str string, cutstr string) string {
// 	return strings.Trim(str, cutstr)
// }
