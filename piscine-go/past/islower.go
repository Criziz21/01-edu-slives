package piscine

func IsLower(s string) bool {
	for _, str := range s {
		if !('a' <= str && str <= 'z') {
			return false
		}
	}
	return true
}