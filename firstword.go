package main

func FirstWord(s string) string {
	str := ""
	for _, i := range s{
		if i == ' '{
			break
		}
		str += string(i)
	}
	return str + "\n"
}
