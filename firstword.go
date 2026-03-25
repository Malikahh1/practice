package main

func FirstWord (s string) string{
	str := ""
	for i := 0; i < len(s); i++{
		if s [i] != ' ' {
			str += string(s[i])
		}else {
			break
		}
	}
	return str
}