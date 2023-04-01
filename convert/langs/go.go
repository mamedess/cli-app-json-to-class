package langs

import (
	"fmt"
	"main/sb"
)

var props = []Prop{}

// decoda o json e retorna uma string representando um struct golang
func CreateGolang() {
	propsMap := decodeJSON([]byte(Str))

	TraverseMap(propsMap, "")
	BuildString()
}

func TraverseMap(m map[string]interface{}, father string) {
	for k, v := range m {
		if isUniqueIn(props, k, k) {

			item := newProp(k, v, "")
			if nestedMap, ok := v.(map[string]interface{}); ok {
				item.father = father
				props = append(props, item)
				TraverseMap(nestedMap, k)
			} else if nestedSlice, ok := v.([]interface{}); ok {
				item.father = father
				props = append(props, item)
				for _, nestedItem := range nestedSlice {
					if nestedMap, ok := nestedItem.(map[string]interface{}); ok {
						TraverseMap(nestedMap, k)
						break
					}
				}
			} else {
				item.father = father
				props = append(props, item)
			}

		}
	}
}

func BuildString() {

	fmt.Println(GenerateStruct(props, false))

	for _, prop := range props {
		if prop.AnyChildren {
			fmt.Println(GenerateStruct(prop.GetChildren(), true))
		}
	}
}

func GenerateStruct(item []Prop, children bool) string {
	if len(item) == 0 {
		return ""
	}

	sb.CreateNew()

	structName := ObjectName

	if children && len(item) > 0 {
		structName = item[0].father
	}

	declaration := fmt.Sprintf("type %v struct {", structName)
	sb.Appen(declaration)

	maxLen := 0

	for _, prop := range item {
		if len(prop.name) > maxLen {
			maxLen = len(prop.name)
		}
	}

	for _, prop := range item {
		if prop.father == "" || children {
			nameWithPadding := fmt.Sprintf("%-"+fmt.Sprintf("%v", maxLen+1)+"s", prop.name)
			sb.Appenl(nameWithPadding)
			sb.Appen(prop.proptype)
		}
	}

	sb.Appenl("}")

	return sb.Retrieve()
}
