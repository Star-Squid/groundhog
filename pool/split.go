package piscine

func Split(s, sep string) []string {
	var result []string
	nextIndex := Index(s, sep)

	for j := 0; nextIndex != -1; j += len(sep) {
		str := s[0:nextIndex]
		result = append(result, str)
		s = s[nextIndex+len(sep):]
		nextIndex = Index(s, sep)
	}

	result = append(result, s)
	return result
}
