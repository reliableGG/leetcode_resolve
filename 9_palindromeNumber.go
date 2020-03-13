func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		p, res := x, 0
        for p != 0  {
			res = res * 10 + p % 10
			p /= 10
        }
		return res == x
	}
}