package piscine

func ToUpper(s string) string {
	str2 := []rune(s)
	for i, str1 := range s {
		if('a' <= str1 && str1 <= 'z') {
			str2[i] = (str1 - 32)
		}
	}
	return string(str2)
}