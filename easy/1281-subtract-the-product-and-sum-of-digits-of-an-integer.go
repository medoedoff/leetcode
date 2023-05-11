package easy

func subtractProductAndSum(n int) int {
	p := product(n)
	s := sum(n)

	return p - s
}

func product(n int) int {
	if n/10 == 0 {
		return n % 10
	} else {
		left := n % 10
		n := n / 10

		return left * product(n)
	}
}

func sum(n int) int {
	if n/10 == 0 {
		return n % 10
	} else {
		left := n % 10
		n := n / 10

		return left + sum(n)
	}
}
