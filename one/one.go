package one

// SumMultiplesOf sums all the natural numbers up to, but not including, below,
// that are multiples of the given multiples
func SumMultiplesOf(below int, multiples ...int) (s int) {
	for i := 1; i < below; i++ {
		if multipleOf(i, multiples) {
			s += i
		}
	}
	return
}

func multipleOf(target int, multiples []int) bool {
	for _, v := range multiples {
		if target%v == 0 {
			return true
		}
	}
	return false
}
