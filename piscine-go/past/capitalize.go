package piscine

// func Capitalize(s string) string {
// 	str2 := []rune(s)
// 	flag := true
// 	for i, str1 := range s {
// 		if((('a' <= str1 && str1 <= 'z') || ('0' <= str1 && str1 <= '9') || ('A' <= str1 && str1 <= 'Z')) && flag) {
// 			if ('0' <= str1 && str1 <= '9') || ('A' <= str1 && str1 <= 'Z') {
// 				flag = false
// 			} else {
// 				str2[i] = (str1 - 32)
// 				flag = false
// 			}
// 		} else if ('A' <= str1 && str1 <= 'Z') {
// 			str2[i] = (str1 + 32)
// 		}
// 		if ((32 <= str1 && str1 <= 47)  || (58 <= str1 && str1 <= 64) || (91 <= str1 && str1 <= 96) || (124 <= str1 && str1 <= 126)) {
// 			flag = true
// 		}
// 	}
// 	return string(str2)
// }

func Capitalize(s string) string {
	IsAlpha := func(r rune) bool {
		return 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || '0' <= r && r <= '9'
	}

	runes := []rune(s)

	for i, char := range runes {
		if i == 0 || !IsAlpha(runes[i-1]) && IsAlpha(char) {
			if 'a' <= char && char <= 'z' {
				runes[i] -= 32 // To upper
			}
		} else {
			if 'A' <= char && char <= 'Z' {
				runes[i] += 32 // To lower
			}
		}
	}

	return string(runes)
}
