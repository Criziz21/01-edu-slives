package piscine

func IsPrintable(s string) bool {
	for _, str := range s {
		if !(32 <= str && str <= 126) {
			return false
		}
	}
	return true
}