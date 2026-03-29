package main

func CamelToSnakeCase3(s string) string {
	if s == "" || !Iscamelcase(s) {
		return s
	}
	result := " "
	for i, r := range s {
		isupper := r >= 'A' && r <= 'Z'

		if i > 0 {
			p := s[i-1]
			precase := p >= 'A' && p <= 'Z'
			if !precase && isupper {
				result += "_"
			}

		}
		result += string(r)
	}
	return result

}

func Iscamelcase1(str string) bool {
	strlen := len(str)

	for i, r := range str {
		Islower := r >= 'a' && r <= 'z'
		Isupper := r >= 'A' && r <= 'Z'

		if i == strlen-1 && Isupper {
			return false
		}
		if i > 0 {
			p := str[i-1]
			prevcase := p >= 'A' && p <= 'Z'
			if prevcase && Isupper {
				return false
			}
		}
		if !Islower && Isupper {
			return false
		}

	}
	return true
}
