package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var wInputMsg string = "wrong input pal, try again!"
var reader = bufio.NewReader(os.Stdin)

// get file from path and return as byte[]
func GetFileInBytes(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	fmt.Println("Contents of file:", string(data))

	return data, err
}

// validate the json inside the start context
func IsValidJson(strJson string) bool {
	isValid := json.Valid([]byte(strings.TrimSpace(strJson)))

	return isValid
}

// called when user input is unknown
func WrongInput() {
	fmt.Println()
	fmt.Println(wInputMsg)
}

// get user prompt and returns it as a string, and error if any
func GetInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

// get regex for speciied pattern
func GetPattern(regex string) *regexp.Regexp {
	return regexp.MustCompile(regex)
}
