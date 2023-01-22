package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var wInputMsg string = "wrong input pal, try again!"
var reader = bufio.NewReader(os.Stdin)
var validTypes = [4]string{"go,golang,c#,csharp"}

// // handle golang choice by validating the json and calling the specified converter function if valid

// func main() {
// 	file, err := os.Open("file.txt")
// 	if err != nil {
// 		fmt.Println("File reading error", err)
// 		return
// 	}
// 	defer file.Close() // it's important to close the file after reading it

//		// create a byte slice to hold the file contents
//		data := make([]byte, 1024)
//		for {
//			n, err := file.Read(data)
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				fmt.Println("File reading error", err)
//				return
//			}
//			fmt.Println("Read", n, "bytes:", string(data[:n]))
//		}
//	}
//
// get file from path and return as byte[]
func GetFileInBytes(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	fmt.Println("Contents of file:", string(data))

	return data, err
}

// validate the json inside the start context
func IsValidJson(jsonPrompt string) bool {
	isValid := !json.Valid([]byte(strings.TrimSpace(jsonPrompt)))
	if isValid {
		fmt.Println("\nwait a second, this is not valid JSON, press 'enter' key to try again")
	}
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
