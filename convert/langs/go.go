package langs

import (
	"fmt"
	"main/sb"
)

var props = []Prop{}
var iteration = 0

// decoda o json e retorna uma string representando um struct golang
func CreateGolang() {
	propsMap := decodeJSON([]byte(Str))
	TraverseMap(propsMap)

	sb.CreateNew()
	declaration := fmt.Sprintf("type %v struct {", ObjectName)
	sb.Appen(declaration)
	for _, prop := range props {
		sb.Appenl(prop.name + "")
		sb.Appen(prop.proptype)
	}
	sb.Appenl("}")

	fmt.Print(sb.Retrieve())
}

func TraverseMap(m map[string]interface{}) {
	for k, v := range m {
		if isUniqueIn(props, k, iteration) {

			item := newProp(k, v)
			if nestedMap, ok := v.(map[string]interface{}); ok {
				iteration++
				item.iteration = iteration

				props = append(props, item)
				TraverseMap(nestedMap)
				iteration = 0
			} else if nestedSlice, ok := v.([]interface{}); ok {
				iteration++
				item.iteration = iteration

				props = append(props, item)
				for _, nestedItem := range nestedSlice {
					if nestedMap, ok := nestedItem.(map[string]interface{}); ok {
						TraverseMap(nestedMap)
						iteration = 0
						break
					}
				}
			} else {
				item.iteration = iteration
				props = append(props, item)
			}

		}
	}
}
