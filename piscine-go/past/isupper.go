package piscine

func IsUpper(s string) bool {
	for _, str := range s {
		if !('A' <= str && str <= 'Z') {
			return false
		}
	}
	return true
}