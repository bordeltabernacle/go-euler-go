package one

func MultiplesOf(below int) (s int) {
	for i := 1; i < below; i++ {
		if i%3 == 0 {
			s += i
		} else if i%5 == 0 {
			s += i
		}
	}
	return
}
