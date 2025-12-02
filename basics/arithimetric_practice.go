package basics

import (
	"fmt"
	"math"
)

func main6() {
	var a, b int = 10, 3
	var result int = a + b
	fmt.Println("Add", result)
	result = a - b
	fmt.Println("Substract", result)
	result = a * b
	fmt.Println("Multiplication", result)
	result = a / b
	fmt.Println("Division", result)
	result = a % b
	fmt.Println("MOD", result)

	const PI float64 = 3.14
	var piResult float64 = 10 / 3.14
	piResult2 := 10 / PI
	fmt.Println(piResult)
	fmt.Println(piResult2)

	var maxInt int = math.MaxInt
	fmt.Println(maxInt)
	maxInt = maxInt + 1
	fmt.Println(maxInt)

	var uMaxInt uint16 = math.MaxUint16
	fmt.Println(uMaxInt)
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var minFloat float64 = 1.0e-323
	fmt.Println(minFloat)
	minFloat = minFloat / math.MaxFloat64
	fmt.Println(minFloat)
}
