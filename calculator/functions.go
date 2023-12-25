package calculator

// gewöhnliches Switch 
func RunPeration(operation string, left, right int) int {
	var result int
	
	switch operation {
		case "add":
			result = left + right
		case "sub":
			result = left - right 
		default:
			//left blank
	}
	return result
}

// switch mit Bedingung im case
func RunPeration(operation string, left, right int) int {
	var result int
	
	switch {
		case operation == "add":
			result = left + right
		case operation == "sub":
			result = left - right 
		default:
			//left blank
	}
	return result
}


if x := 100; x == 100 {
		fmt.Println("Germany")
}

