// first idea: using map to filter accept string
func countCharacters(words []string, chars string) int {
	wordMap := make(map[string]int)
	res := 0
	for i := 0; i < len(chars); i++ {
		wordMap[string(chars[i])]++
	}

	for i := 0; i < len(words); i++ {
		word := words[i]
		wMap := make(map[string]int)
		for j := 0; j < len(word); j++ {
			wMap[string(word[j])]++
		}

		judge := true
		for key, value := range wMap {
			if wordMap[key] < value {
				judge = false
			}
		}

		if judge {
			res += len(words[i])
		}
	}

	return res
}

// fatest solution
func countCharacters(words []string, chars string) int {
    count := make([]int, 26)
    for i := 0; i < len(chars); i++ {
        count[chars[i]-'a']++
    }
    ans := 0
    for _, w := range words {
        if canBeFormed(w, count) {
            ans += len(w)
        }
    }
    return ans
}
func canBeFormed(w string, c []int) bool {
    count := make([]int, 26)
    for i := 0; i < len(w); i++ {
        count[w[i]-'a']++
        if count[w[i]-'a'] > c[w[i]-'a'] {
            return false
        }
    }
    return true
}