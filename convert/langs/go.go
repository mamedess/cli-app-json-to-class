package langs

import (
	"fmt"
)

func CreateGolang() {
	data := []byte(Str)
	n, err := decodeJSON(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}
