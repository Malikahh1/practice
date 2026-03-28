package main


func CountAlpha(str string) int {
	Count := 0
	for _, i:= range str {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			Count++
		}
	}
	return Count 
}

