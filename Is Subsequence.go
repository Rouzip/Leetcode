func isSubsequence(s string, t string) bool {
	lengthS := len(s)
	lengthT := len(t)

	if lengthS == 0 {
		return true
	}

	checkArr := make([]bool, lengthS)

	sIndex := 0
	for tIndex := 0; tIndex < lengthT; tIndex++ {
		if s[sIndex] == t[tIndex] {
			checkArr[sIndex] = true
			sIndex++
			if sIndex == lengthS {
				break
			}
		}
	}
	return checkArr[lengthS-1]
}