package main



func Retainfirsthalf(str string) string {
	if str == "" {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	halflength := len(str)/ 2
	return str[:halflength]

}