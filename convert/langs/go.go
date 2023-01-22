package langs

import "fmt"

func ConvertGolang() {
	n := GetPropNames(Str)
	uniqueNames := []string{}
	l := ""

	for i := 0; i < len(n); i++ {
		cprop := n[i]
		if l != "" && cprop != l {
			l = cprop
			uniqueNames = append(uniqueNames, cprop)
		}
	}

	fmt.Println("\nfound these:")
	fmt.Println(uniqueNames)
}
