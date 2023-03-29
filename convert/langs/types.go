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
	father      string
	AnyChildren bool
	GetChildren func() []Prop
}

func newProp(name string, item interface{}, father string) Prop {
	return Prop{
		name:        name,
		proptype:    RetrieveType(item),
		father:      father,
		AnyChildren: strEquals(RetrieveType(item), "map[string]interface{}") || strEquals(RetrieveType(item), "[]interface{}"),
		GetChildren: func() []Prop {
			var children = []Prop{}
			if nestedMap, ok := item.(map[string]interface{}); ok {
				for k, v := range nestedMap {
					if isUniqueIn(children, k, name) {
						children = append(children, newProp(k, v, name))
					}
				}
			}

			if nestedSlice, ok := item.([]interface{}); ok {
				for _, prop := range nestedSlice {
					if nested, ok := prop.(map[string]interface{}); ok {
						for k, v := range nested {
							if isUniqueIn(children, k, name) {
								children = append(children, newProp(k, v, name))
							}
						}
					}
				}
			}
			return children
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
