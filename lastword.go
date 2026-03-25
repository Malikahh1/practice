package main

func LastWord(s string) string {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	j := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return s[i+1:j+1] + "\n"
}
