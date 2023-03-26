package langs

type PropNValue struct {
	name         string
	value        interface{}
	GetValueType func() string
	AnyChildren  bool
	GetChildren  func() map[string]interface{}
	GetProp      func() Prop
}

type Prop struct {
	name        string
	proptype    string
	iteration   int
	AnyChildren bool
}

func newProp(name string, item interface{}) Prop {
	return Prop{
		name:        name,
		proptype:    RetrieveType(item),
		AnyChildren: strEquals(RetrieveType(item), "map[string]interface{}") || strEquals(RetrieveType(item), "[]interface{}"),
	}
}

func newPropNValue(name string, item interface{}) PropNValue {
	return PropNValue{
		name:  name,
		value: item,
		GetValueType: func() string {
			return RetrieveType(item)
		},
		GetProp: func() Prop {
			return newProp(name, RetrieveType(item))
		},
		AnyChildren: strEquals(RetrieveType(item), "map[string]interface{}") || strEquals(RetrieveType(item), "[]interface{}"),
		GetChildren: func() map[string]interface{} {
			return item.(map[string]interface{})
		},
	}
}

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
	case float64:
		assertedtype = "float64"
	case []interface{}:
		assertedtype = "[]interface{}"
	case interface{}:
		assertedtype = "interface{}"
	}
	return assertedtype
}
