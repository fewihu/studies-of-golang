package composite

func DemoCollections() {
	primesArray := int[5]{2,3,5,7,11}

	primesSlice := int[]{2,3,5,7,11}
	fmt.Println(len(primesSlice))
	fmt.Println(cap(primesSlice))


	randomSlice := make([]int,5,7)
	randomSlice  = append(randomSlice, 42)

	for index, value := range primesSlice {
		fmt.Println(index, value)
	}

	for _, value := range primesSlice {
		fmt.Println(value)
	}

	for index, _ := range primesSlice {
		fmt.Println(index)
	}

	points :=map[string]point{
		"A": *NewPoint(2,3),
		"B": *NewPoint(4,5)
	}


	//value a, true
	pointA , ok := point["A"]
	fmt.Println(pointA, ok)

	//defaults, false
	pointA , ok := point["D"]
	fmt.Println(pointA, ok)


	points["E"] = *NewPoint(1,1)
	delete(points, "A")
}
