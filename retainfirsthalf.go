package main

func Retainfirsthalf(str string) string {
	if len(str) == 0 {
		return str
	}
	if len(str) == 1 {
		return str
	}

	half := len(str) / 2
	return str [:half]

}

// func main () {
// 	if len(str) <= 1 {
// 		return str
// 	}

// 	return str[:len(str)/2]
// }
