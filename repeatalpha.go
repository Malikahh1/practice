package main

func RepeatAlpha (s string) string {
	result := ""
	count := 1

	for _, r range s{
		if >= 'a' && r <= 'z'{
			count = int(r-a+1)
		}
		if >= 'A' && r <= 'Z'{
			count = int(r-A+1)
		}
		for i :=0; i < count; i++{
			result +=string(r)
		}
	}
	return result
	
}
