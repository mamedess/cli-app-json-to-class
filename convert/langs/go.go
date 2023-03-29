package langs

import (
	"fmt"
	"main/sb"
)

var props = []Prop{}
var secondaryProps = []Prop{}

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
	sb.CreateNew()
	declaration := fmt.Sprintf("type %v struct {", ObjectName)
	sb.Appen(declaration)

	maxLen := 0

	for _, prop := range props {
		if len(prop.name) > maxLen {
			maxLen = len(prop.name)
		}
	}

	for _, prop := range props {
		if prop.father == "" {
			nameWithPadding := fmt.Sprintf("%-"+fmt.Sprintf("%v", maxLen+1)+"s", prop.name)
			sb.Appenl(nameWithPadding)
			sb.Appen(prop.proptype)
		} else {
			if isUniqueIn(secondaryProps, prop.father, "") && prop.proptype == "[]interface{}" || prop.proptype == "map[string]interface{}" {
				secondaryProps = append(secondaryProps, prop.GetChildren()...)
			}
		}
	}

	sb.Appenl("}")

	fmt.Print(sb.Retrieve())
	var t = secondaryProps
	fmt.Print(t)
}
