package langs

func RetrieveType(value interface{}) string {
	assertedtype := ""

	switch value.(type) {
	case map[string]interface{}:
		assertedtype = "map[string]interface{}"
	case string:
		assertedtype = "string"
	case []string:
		assertedtype = "string[]"
	case int64:
		assertedtype = "int"
	case int:
		assertedtype = "int"
	case float64:
		assertedtype = "float64"
	case []interface{}:
		assertedtype = "[]interface{}"
	case interface{}:
		assertedtype = "interface{}"
	}
	return assertedtype
}
