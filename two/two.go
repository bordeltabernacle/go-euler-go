package two

// SumEvenFib takes the terms in the Fibonacci sequence whose values do not
// exceed n and finds the sum of the even-valued terms.
func SumEvenFib(n int) (s int) {
	for a, b := 0, 1; a < n; a, b = a+b, a {
		if a%2 == 0 {
			s += a
		}
	}
	return
}
