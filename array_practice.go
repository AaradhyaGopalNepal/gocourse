package main

import "fmt"

func main() {
	var array [10]int
	fmt.Println(array)
	array[2] = 100
	fmt.Println(array)

	fruits := [4]string{"Apple", "Orange", "Banana", "Mango"}
	fmt.Println(fruits)

	originalArray := [4]int{1, 2, 3, 4}
	copiedArray := originalArray
	copiedArray[2] = 100
	fmt.Println(originalArray)
	fmt.Println(copiedArray)

	for index, value := range originalArray {
		fmt.Print(index+1, " Apple ")
		for range value {
			fmt.Print("Japak ")
		}
		fmt.Println()
	}
}
