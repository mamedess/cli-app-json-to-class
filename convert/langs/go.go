package langs

import (
	"fmt"
	"main/sb"
	"os"
)

func CreateGolang() {
	n := GetPropNames(Str)

	sb.CreateNew()
	sb.Appen(fmt.Sprintf("type %v struct {", Name))

	for i := 0; i < len(n); i++ {
		c := n[i]
		sb.Appenl(fmt.Sprintf("%4v %v", c.name, GetValueType(c.value)))
	}

	sb.Appenl("}")

	data := []byte(sb.Retrieve())
	filename := Name + ".go"
	err := os.WriteFile("output/"+filename, data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("\nfile was save as ", Name+".go", "in output folder")
}
