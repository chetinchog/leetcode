package hackerrank

import "regexp"

// BracesValidation returns array of validated results.
func BracesValidation(values []string) []string {
	var result []string
	for _, list := range values {
		status := "YES"
		var temp string
		for len(list) > 0 {
			char, listL := divideFirst(list)
			list = listL
			ok,e:=regexp.MatchString("[\(\[\{]", char)
			if ok {
				temp += char
			} else {
				switch char {
				case ")":
					break
				case "]":
					break
				case "}":
					break
	
				}
			}
		}
		result = append(result, status)
	}
	return result
}

func divideFirst(values string) (first string, rest string) {
	first, rest = string(values[0]), values[1:]
	return
}

func divideLast(values string) (rest string, last string) {
	rest, last = values[:len(values)-2], string(values[len(values)-1])
	return
}
