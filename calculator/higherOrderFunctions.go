package calculator

func SumFromAtoB(a, b int) int {
	return ProcessFromAtoB(a, b, 0, func(x, y int) int {
		return x + y
	})
}

func MultiplyFromAtoB(a,b int) int {
	return ProcessFromAtoB(a, b, 1, func(x, y int) int) {
		return x * y;
	}
}

func ProcessFromAtoB(a,b, initValue int, fn func(int, int) int) int {
	increment := AddX(1)
	
	if a > b {
		return initValue
	}

	return fn(a, ProcessFromAtoB(increment(a), b, initValue, fn))
} 

func AddX(x int) func(int) int {
	return func(y int) int {
		x + y
	}
}  
