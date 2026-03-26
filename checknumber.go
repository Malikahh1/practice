package main

func CheckNumber(s string) bool {
	for i := range s {
		if i >= '0' && i <= '9' {
			return true
		}
	}
	return false
}
