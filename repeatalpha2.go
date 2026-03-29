package main


func RepeatAlpha2(s string) string {
	result := ""

	for _, r := range s {
		count := 1
		if r >= 'a' && r <= 'z'{
			count = int(r-'a' + 1)
		}
		if r >= 'A' && r <= 'Z' {
			count = int(r - 'A' + 1)
		}
		for i := 0; i < count; i++{
			result += string(r)
		}
		
	}
	return result
}
