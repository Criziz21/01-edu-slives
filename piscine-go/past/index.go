package piscine

func Index(s, toFind string) int {
	str1, str2 := []rune(s), []rune(toFind)
	if len(str2) == 0 || len(str1) == 0 {
		return 0
	}
	if len(str1) < len(str1) {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if str1[i] == str2[0] {
			if len(str2) == 1 {
				return i
			} else {
				for j := 0; j < len(toFind); j++ {
					if str1[i+j] != str2[j] {
						break
					} 
					if j == (len(toFind) - 1) {
						return i
					}
				}
			}
		}
	}
	return -1
}
