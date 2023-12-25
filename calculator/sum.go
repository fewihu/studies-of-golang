package calculator

func NearestBinaryPower (val int) int {
	sum := 1
	
	for sum < val {
		sum += sum
	}
	
	return sum
}

func Sum (min, max int) int {
	sum := 0
	
	for i:=0; i<max; i++ {
		sum += i
	}
	
	return sum
}
