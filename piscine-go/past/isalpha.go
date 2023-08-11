package piscine

func IsAlpha(s string) bool {
	for _, str := range s {
		if !(('a' <= str && str <= 'z') || ('A' <= str && str <= 'Z') || ('0' <= str && str <= '9') || ' ' == str) {
			return false
		}
	}
	return true
}
