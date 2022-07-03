package piscine

func MaxWordCountN(text string, n int) map[string]int {
	words := split(text)
	wordFrequency := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		freq := 1

		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			} else {
				if words[i] == words[j] {
					freq++
				}
			}
		}
		wordFrequency[i] = freq
	}

	allWords := make(map[string]int)

	for i := 0; i < len(words); i++ {
		allWords[words[i]] = wordFrequency[i]
	}

	// create struct for efficient sorting
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for Key, Value := range allWords {
		ss = append(ss, kv{Key, Value})
	}

	// sort by value
	for i := 0; i < len(ss)-1; i++ {
		for j := 0; j < len(ss)-i-1; j++ {
			if ss[j].Value < ss[j+1].Value {
				ss[j], ss[j+1] = ss[j+1], ss[j]
			}
		}
	}

	// sort by key if values are the same
	for i := 0; i < len(ss)-1; i++ {
		for j := 0; j < len(ss)-i-1; j++ {
			if (ss[j].Value == ss[j+1].Value) && (ss[j].Key > ss[j+1].Key) {
				ss[j], ss[j+1] = ss[j+1], ss[j]
			}
		}
	}

	// add n values
	result := make(map[string]int, n-1)
	for _, kv := range ss[:n] {
		result[kv.Key] = kv.Value
	}

	return result
}

func split(s string) []string {
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

	return result
}
