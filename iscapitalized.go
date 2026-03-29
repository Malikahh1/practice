package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	// 1. Check if the string is empty
	if s == "" {
		return false
	}
	
	// 2. Check if the first character is a lowercase letter
	for i, r := range s{
		Islower := r >= 'a' &&   r <= 'z'
		if i ==  0 && Islower {
			return false
		}
		if i > 0 && s[i-1] == ' '&& Islower{
		return false
	}
}
return true
}
