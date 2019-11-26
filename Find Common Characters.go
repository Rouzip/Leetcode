// first idea, use map to get right answer
// append availble letters to the array 
func commonChars(A []string) []string {
	length := len(A)
	lMaps := make([]map[byte]int, length)
	for i := 0; i < length; i++ {
		lMaps[i] = make(map[byte]int)
	}

	leastIndex := 0
	max := 100
	for i := 0; i < length; i++ {
		str := A[i]
		strLength := len(str)
		if max > strLength {
			max = strLength
			leastIndex = i
		}

		for j := 0; j < strLength; j++ {
			lMaps[i][str[j]]++
		}
	}

	res := []string{}

	for key := range lMaps[leastIndex] {
		v := 101
		fmt.Println(string(key))

		for i := 0; i < length; i++ {
			curStringVal, _ := lMaps[i][key]
			if curStringVal < v {
				v = curStringVal
			}
		}

		if v != 101 {
			for i := 0; i < v; i++ {
				res = append(res, string(key))
			}
		}
	}

	return res
}

// the fastest idea
func commonChars(A []string) []string {
    var mapper [26]int
    tmp := [26]int{}

    for i := 0; i < 26; i ++ {
        mapper[i] = 100
    }

    for _, str := range A {
        for _, ch := range []byte(str) {
            tmp[ch - 'a'] ++
        }
        for i := 0; i < 26; i ++ {
            if tmp[i] < mapper[i] {
                mapper[i] = tmp[i]
            }
            tmp[i] = 0
        }
    }

    output := []string{}
    for i := 0; i < 26; i ++ {
        ch := string([]byte{byte(i) + 'a'})
        for j := 0; j < mapper[i]; j ++ {
            output = append(output, ch)
        }
    }

    return output
}

// use least memory
func commonChars(A []string) []string {
	com := make(map[string]int)
	word := A[0]
	chars := []byte(word)
	for _, c := range chars {
		com[string(c)] = strings.Count(word, string(c))
	}
	for i := 1; i < len(A); i++ {
		for ch, cnt := range com {
			n := strings.Count(A[i], ch)
			if n == 0 {
				delete(com, ch)
			} else if n > 0 && n < cnt {
				com[ch] = n
			}
		}
	}
	comc := make([]string, 0)
	for k, v := range com {
		for i := 0; i < v; i++ {
			comc = append(comc, k)
		}
	}
	return comc    
}