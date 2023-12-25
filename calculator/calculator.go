package calculator

func Add(left, right int) int {
	return left + right
}

func Divide(divident, divisor int) (quotient int, remainder int) {
	quotient = divident / divisor
	remainder = divident % divisor
	return quotient, remainder
}
