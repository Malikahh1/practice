package 

func CamelToSnakeCase(s string) string {
	if s == "" || !Iscamelcase(s) {
		return s
	}

	result := ""
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

func Iscamelcase(str string) bool {
	strlen := len(str)

	for i, r := range str {
		isupper := r >= 'A' && r <= 'Z'
		islower := r >= 'a' && r <= 'z'

		if i == strlen-1 && isupper {
			return false
		}
		if i > 0 {
			p := str[i-1]
			precase := p >= 'A' && p <= 'Z'
			if isupper && precase {
				return false
			}
		}
		if !isupper && !islower {
			return false
		}
	}
	return true
}
