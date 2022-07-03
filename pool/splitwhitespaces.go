package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)
	var str string
	var result []string

	for i := 0; i < len(runes); i++ {
		if runes[i] == 32 || runes[i] == 9 || runes[i] == 10 {
			result = append(result, str)
			str = ""
		} else {
			str = str + string(runes[i])
		}
	}
	result = append(result, str)

	// in case there was a double space that created an empty string
	// var result2 []string
	// if len(result) > 0 {
	// 	for j := 0; j < len(result); j++ {
	// 		if result[j] != "" {
	// 			result2 = append(result2, result[j])
	// 		}
	// 	}
	// }
	return result
}
