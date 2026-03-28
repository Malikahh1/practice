package main

import "strconv"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	step := 1
	result := ""
	if from > to {
		step = -1
	}

	for i := from; ; i += step {
		numstr := strconv.Itoa(i)
		if i < 10 {
			numstr = "0" + numstr
		}
		result += numstr
		if i == to {
			break
		}
		result += ", "
	}
	return result + "\n"
}
