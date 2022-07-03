package piscine

func ConcatParams(args []string) string {
	var empty string
	if len(args) > 0 {
		result := args[0]
		for i := 1; i < len(args); i++ {
			result = result + "\n" + args[i]
		}
		return result
	}
	return empty
}
