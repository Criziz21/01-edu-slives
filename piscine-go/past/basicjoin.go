package piscine

func BasicJoin(elems []string) string {
	base := ""
	for _, el := range elems {
		base += el
	}
	return base
}