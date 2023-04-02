package langs

import (
	"encoding/json"
)

var Str, ObjectName = "", ""

// Decode the json to a map of strings and its values as interface
func decodeJSON(data []byte) []map[string]interface{} {
	var jsonData []map[string]interface{}
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		panic(err)
	}

	return jsonData
}

// Checks if the current item is already in the specified []
func isUniqueInMap(slice []map[string]interface{}, name string, value interface{}) bool {
	for i := 0; i < len(slice); i++ {
		for k, val := range slice[i] {
			if k == name && RetrieveType(val) == RetrieveType(value) && val == value {
				return false
			}
		}
	}

	return true
}
