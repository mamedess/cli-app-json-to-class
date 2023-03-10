package langs

import (
	"encoding/json"
	"errors"
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

func isUnique(name string, slice []PropNValue) bool {
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

func GetProp(currentprop PropNValue) Prop {
	ptype, err := AssertType(currentprop)

	if err != nil {
		panic(err)
	}

	return newProp(currentprop.name, ptype)
}

// retorna uma instancia de Prop após asserção de tipos
func AssertType(currentprop PropNValue) (string, error) {
	assertedtype := ""

	switch currentprop.value[0].(type) {
	case map[string]interface{}:
		assertedtype = "map[string]interface{}"
	case string:
		assertedtype = "string"
	case []string:
		assertedtype = "string[]"
	case int64:
		assertedtype = "int"
	case float64:
		assertedtype = "float64"
	case interface{}:
		assertedtype = "interface{}"
	default:
		{
			return assertedtype, errors.New("propriedade com tipo desconhecido")
		}
	}
	return assertedtype, nil
}
