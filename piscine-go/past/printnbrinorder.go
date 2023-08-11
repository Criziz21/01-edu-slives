package piscine

const (
	MaxInt32 = 1<<31 - 1
	MaxInt64 = 1<<63 - 1
)

func atoi(s string) int {
	var n, f int64
	n = 0
	f = 1

	minus := 0
	plus := 0
	for i := 0; i < len(s); i++ {
		b := int(s[i])
		if !(b >= 48 && b <= 57 || b == 43 || b == 45) {
			return 0
		}
		if (b == 45 || b == 43) && i > 0 {
			return 0
		}
		if b == 43 {
			plus++
		}

		if b == 45 {
			minus++
		}

		if plus > 1 || minus > 1 {
			return 0
		}
	}

	count := plus + minus
	s = s[count:]
	for j := 1; j < len(s); j++ {
		f *= 10
	}

	for i := 0; i < len(s); i++ {
		num := int64(int(s[i]) - '0')
		n += num * f
		f /= 10
	}

	if n > MaxInt64 {
		return 0
	}

	if minus > 0 {
		n *= -1
	}
	return int(n)
}

func TrimAtoi(s string) int {
	str1 := []rune(s)
	str2 := []rune{}
	for i := 0; i < len(str1); i++ {
		if ('0' <= str1[i] && str1[i] <= '9') || str1[i] == '-' {
			if str1[i] == '-' {
				if len(str2) == 0 {
					str2 = append(str2, str1[i])
				}
			} else {
				str2 = append(str2, str1[i])
			}
		}
	}
	return atoi(string(str2))
}