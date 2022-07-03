package piscine

func Index(s string, toFind string) int {
	// cases depend on respective lengths of strings
	switch {
	case len(toFind) == 0:
		return 0

	case len(toFind) == 1:
		// return its index
		for i := 0; i <= len(s); i++ {
			if s[i] == toFind[0] {
				return i
			}
		}

	case len(toFind) == len(s):
		if toFind == s {
			return 0
		}
		return -1

	case len(toFind) > len(s):
		return -1

	case len(toFind) < len(s):
		// find start of toFind in s
		var startingIndex int

		for i := 0; i <= len(s)-1; i++ {
			if s[i] == toFind[0] {
				startingIndex = i
				// check if section of s starting at startingIndex and ending at startingIndex + (n-1) equals toFind
				// if yes, return startingIndex
				// if no, keep trying
				// if no match, return -1
				if string(s[startingIndex:startingIndex+len(toFind)]) == toFind {
					return startingIndex
				}
			}
		}

	}
	return -1
}
