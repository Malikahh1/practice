package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	result := ""

	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	if result == "" {
		return "\n"
	}

	return result + "\n"
}

func main() {
	fmt.Println(ThirdTimeIsACharm("123456789"))
}
