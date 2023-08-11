package piscine

func Join(elems []string, sep string) string {
	base := elems[0]
	for i := 1; i < len(elems); i++ {
		base += (sep + elems[i])
	}
	return base
}