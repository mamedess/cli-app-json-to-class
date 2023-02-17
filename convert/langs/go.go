package langs

import "fmt"

func CreateGolang() {
	//decoda a string em um map[string]interface{}
	props := decodeJSON([]byte(Str))
	//p armazena todas os valores unicos deste json, e tambem seus respectivos tipos
	p := []PropNValue{}

	//chama a funcao is isUnique em cada item de props e armazena valores unicos em p
	for k, value := range props {
		if isUnique(k, p) {

			item := PropNValue{
				name:  k,
				value: []interface{}{value},
			}

			p = append(p, item)
		}
	}
	var pt []Prop
	for i := 0; i < len(p); i++ {
		pt = append(pt, AssertType(p[i]))
	}

	fmt.Println(pt)
}
