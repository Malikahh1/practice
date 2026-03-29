package main

import "fmt"

func CheckNumber (s string) bool {
	for _, i :=  range s {
		if i >= '0' && i <= '9' {
			return true
		}
		

	}
	return false
}

func mainan () {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello0"))
}