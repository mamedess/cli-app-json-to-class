package langs

import (
	"fmt"
)

func CreateGolang() {
	props := decodeJSON([]byte(Str))
	p := []Prop{}

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

		switch t := c.(type) {
		case map[string]interface{}:
			{
				for kkk, vvv := range t {
					fmt.Println(kkk, vvv, "INNER")
				}
			}
		case string:
			fmt.Println(t, "is an string")
		case []string:
			fmt.Println(t, "is a string[]")
		case int64:
			fmt.Println(t, "is a int")
		case float64:
			fmt.Println(t, "is a float")
		case interface{}:
			{
				slice, ok := c.([]interface{})
				if ok {
					for kkk, vvv := range slice {
						tes, okk := vvv.(map[string]interface{})
						if okk {
							for kkkk, vvvv := range tes {

								fmt.Println(kkk, vvv, kkkk, vvvv, "INNNNNER")
							}
						}
					}
				}
			}
		default:
			{
				panic("aaaaa")
			}
		}
	}
}
