package main

func FirstWord(s string) string {
	str := ""
	for _, r := range s {
		if r == ' ' {
			break
		}
		str +=string(r)
	}
	return str + "\n"
}
