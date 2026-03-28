package main

import "fmt"

func Digilen(n, base int)int{
	 if base < 2 || base > 36 {
		return -1
	 }
	 if n < 0 {
		n = -n
	 }
	 counter := 0
	 for n > 0 {
		n/=base
		counter++
	 }
	 return counter 
}

func maisn () {
	fmt.Println(Digilen(-100,10))
}