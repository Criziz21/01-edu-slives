package piscine

func IsNumeric(s string) bool {
	for _, str := range s {
		if !('0' <= str && str <= '9') {
			return false
		}
	}
	return true
}
