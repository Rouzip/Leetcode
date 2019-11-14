// math idea, if the first passenger cannot get his seat, other passengers 
// has same probability to take his own seat
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 1.0 / 2.0
}

// recursive method
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 1.0/float64(n) + float64(n-2)/float64(n)*nthPersonGetsNthSeat(n-1)
}