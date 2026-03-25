package main

import "fmt"

func FishAndChips(n int )string{
	if n%2 == 0{
		return "Fish"
	}
	if n%3 == 0 {
		return "chips"
	}
	if (n%2 == 0) && (n%3 == 0){
		return "Fish and chips"
	}
	if n < 0 {
		return "error: number is negative"
	}
	return "error: non divisible"
}

func maibn (){
	fmt.Println(FishAndChips(-1))
}