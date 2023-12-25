package main

import (
	"fmt"
	"fm.de/hallowelt/calculator"
)

func main() {
	fmt.Println("23 + 42 =", calculator.Add(23,42))
	fmt.Println("17 / 3  =")
	fmt.Println(calculator.Divide(17,3))
	fmt.Println("Summe 1 bis 9:", calculator.Sum(1,10))
	val := 1000
	fmt.Println("Nächste 2er Potenz von", val, calculator.NearestBinaryPower(val))
	
	point := composite.Point{}
	
	point := composite.Point{3, 7}
	
	point := composite.Point{
		X: 3, 
		Y: 7
	}
	
	fmt.Println(point.X, point.Y)
	
	pointPtr := &point
	fmt.Println(pointPtr.X, pointPtr.Y)
}
