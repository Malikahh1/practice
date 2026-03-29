package main

func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if Isprime(i) {
			return i
		}
	}
	return 0
}

func Isprime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
