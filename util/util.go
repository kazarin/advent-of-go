package util

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func Cycle(slice string) func() byte {
	i := 0
	return func() byte {
		value := slice[i]
		i = (i + 1) % len(slice)
		return value
	}

}
