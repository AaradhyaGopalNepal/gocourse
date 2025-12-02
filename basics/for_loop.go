package basics

import "fmt"

func main4() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
	}

	numbers := []int{1, 2, 3}
	for index, array := range numbers {
		fmt.Println("In", index, "there is", array)
	}

	for i := range 10 {
		fmt.Println(i)
	}

}
