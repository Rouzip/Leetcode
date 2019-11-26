func reverseString(s []byte)  {
	slow, quick := 0, len(s) - 1
	
	for slow <= quick {
		s[slow], s[quick] = s[quick], s[slow]
		slow++
		quick--
	}
}


// fastest solution
func reverseString(s []byte)  {
    l,r := 0, len(s)-1
    
    for l < r {
        s[l], s[r] = s[r], s[l]
        l,r = l+1, r-1
    }
}
