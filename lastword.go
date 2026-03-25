package main

func LastWord (s string) string {
	count := 0

	for i := 0; i < len(s); i++{
		if s[i] != ' '{
			count ++
		} else {
			break
		}
	}
	return s[count+1:]
}