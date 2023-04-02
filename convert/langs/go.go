package langs

import (
	"fmt"
	"main/sb"
	"strings"
	"time"
)

var props = make([]map[string]interface{}, 1)

// Decodes the json into a map[string]interface{} and
// traverse its map generatin a []map[string]interface{}.
// Each position of the array corresponds to a struct to be created.
// Finally prints a Golang struct in the prompt.
func CreateGolang() {
	if !strings.HasPrefix(Str, "[") && !strings.HasSuffix(Str, "]") {
		Str = "[" + Str[:1] + Str[1:] + "]"
	}

	propsMap := decodeJSON([]byte(Str))

	start := time.Now()
	for i := 0; i < 1; i++ {

		props = make([]map[string]interface{}, 1)
		props[0] = make(map[string]interface{})
		TraverseMap(propsMap[0], 0, "", nil)
		GenerateStruct(props)
	}

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

				kf = ""
				vf = nil
			}

		}
	}
}

// Generate a struct for each item of the array and print it to the prompt.
func GenerateStruct(items []map[string]interface{}) {
	iteration := 0
	sb.CreateNew()

	for _, item := range items {
		maxLen := 0

		for k := range item {
			kLen := len(k)
			if kLen > maxLen {
				maxLen = kLen
			}
		}

		structName := ObjectName

		if iteration > 0 {
			for k, prop := range item {
				typp := RetrieveType(prop)

				if validType(typp) {
					structName = k
					delete(item, k)
					break
				}

			}
		}

		sb.AppenEmptyl()

		declaration := fmt.Sprintf("type %v struct {", structName)
		sb.Appen(declaration)

		for k, prop := range item {
			nameWithPadding := fmt.Sprintf("%-"+fmt.Sprintf("%v", maxLen+1)+"s", k)
			sb.Appenl(nameWithPadding)
			sb.Appen(RetrieveType(prop))

		}
		iteration++
		sb.Appenl("}")
		sb.AppenEmptyl()
	}

	fmt.Print(sb.Retrieve())
}

// Checks if type is a map of strings or slice interface
func validType(str string) bool {

	for _, s := range [2]string{"map[string]interface{}", "[]interface{}"} {
		if s == str {
			return true
		}
	}

	return false
}
