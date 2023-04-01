package langs

import (
	"fmt"
	"main/sb"
	"time"
)

var props = make([]map[string]interface{}, 1)

// Decodes the json into a map[string]interface{} and
// traverse its map generatin a []map[string]interface{}.
// Each position of the array corresponds to a struct to be created.
// Finally prints a Golang struct in the prompt.
func CreateGolang() {
	propsMap := decodeJSON([]byte(Str))

	start := time.Now()

	props = make([]map[string]interface{}, 1)
	props[0] = make(map[string]interface{})
	TraverseMap(propsMap, 0, "", nil)
	GenerateStruct(props)

	elapsed := time.Since(start)
	fmt.Printf("it took %s", elapsed)
}

// Traverse a map[string]interface{} using recursion.
func TraverseMap(m map[string]interface{}, iteration int, kf string, vf interface{}) {
	for k, v := range m {
		if isUniqueInMap(props, k, v) {

			if kf != "" && vf != nil {
				props = append(props, make(map[string]interface{}))
				iteration = len(props) - 1
				props[iteration][kf] = vf
			}

			if nestedMap, ok := v.(map[string]interface{}); ok {
				props[iteration][k] = v
				TraverseMap(nestedMap, iteration, k, v)
			} else if nestedSlice, ok := v.([]interface{}); ok {
				props[iteration][k] = v
				for _, nestedItem := range nestedSlice {
					if nestedMap, ok := nestedItem.(map[string]interface{}); ok {
						TraverseMap(nestedMap, iteration, k, v)
						break
					}
				}
			} else {
				props[iteration][k] = v

				if kf != "" && vf != nil {
					kf = ""
					vf = nil
				}
			}

		}
	}
}

// Generate a struct for each item of the array and print it to the prompt.
func GenerateStruct(items []map[string]interface{}) {
	sb.CreateNew()

	for _, item := range items {

		structName := "teste"

		declaration := fmt.Sprintf("type %v struct {", structName)
		sb.Appen(declaration)

		maxLen := 0

		for k := range item {
			if len(k) > maxLen {
				maxLen = len(k)
			}
		}

		for k, prop := range item {
			nameWithPadding := fmt.Sprintf("%-"+fmt.Sprintf("%v", maxLen+1)+"s", k)
			sb.Appenl(nameWithPadding)
			sb.Appen(RetrieveType(prop))
		}

		sb.Appenl("}")
		sb.Appenl("")

	}

	fmt.Print(sb.Retrieve())
}
