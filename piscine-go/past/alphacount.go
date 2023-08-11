package piscine

func AlphaCount(s string) int {
	cnt := 0
	for _, str := range s {
		if(('a' <= str && str <= 'z') || ('A' <= str && str <= 'Z')) {
			cnt++
		}
	}
	return cnt
}