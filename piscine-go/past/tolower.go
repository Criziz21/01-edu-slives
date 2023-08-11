package piscine

func ToLower(s string) string {
	str2 := []rune(s)
	for i, str1 := range s {
		if('A' <= str1 && str1 <= 'Z') {
			str2[i] = (str1 + 32)
		}
	}
	return string(str2)
}