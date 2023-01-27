package langs

import (
	"fmt"
	"main/sb"
	"reflect"
)

func CreateGolang() {
	props := decodeJSON([]byte(Str))
	p := []Prop{}

	sb.CreateNew()
	sb.Appen(fmt.Sprintf("type %v struct {", Name))

	for k, value := range props {
		if isUnique(k, p) {

			item := Prop{
				name:  k,
				value: []interface{}{value},
			}

			p = append(p, item)
		}
	}

	for i := 0; i < len(p); i++ {
		c := p[i].value[0]

		rt := reflect.TypeOf(c)
		switch rt.Kind() {
		case reflect.Slice:
			fmt.Println(c, "is a slice")
		case reflect.Array:
			fmt.Println(c, "is an array")
		case reflect.String:
			fmt.Println(c, "is a string")
		case reflect.Int:
			fmt.Println(c, "is a int")
		case reflect.Float64:
			fmt.Println(c, "is a float")
		case reflect.Interface:
			fmt.Println(c, "is a interface")
		default:
			{
				fmt.Println(c, "is something else entirely")
				slice, ok := c.(map[string]interface{})
				if ok {
					for kkk, vvv := range slice {
						fmt.Println(kkk, vvv)
					}
				} else {
					panic("aaaaa")
				}
			}
		}
	}

	// fmt.Println(p)
	// data := []byte(sb.Retrieve())
	// filename := Name + ".go"
	// err := os.WriteFile("output/"+filename, data, 0644)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("\nfile was save as ", Name+".txt", "in output folder")
}
