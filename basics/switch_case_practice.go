package main

import "fmt"

func main() {
	checkType(1)
	checkType(1.1)
	checkType("Hello")
	checkType(false)
	fruit := "Apple"
	switch fruit {
	case "Apple":
		fmt.Println("Red")
	case "Banana":
		fmt.Println("Yello")
	default:
		fmt.Println("No Idea")
	}

	dayOfWeek := "tuesday"
	switch dayOfWeek {
	case "saturday", "sunday":
		fmt.Println("WeekEnd")
	case "monday":
		fmt.Println("Boring Day")
	case "tuesday", "wednesday", "thrusday":
		fmt.Println("Work Day")
	case "friday":
		fmt.Println("No work, specially no deploy to production day")
	default:
		fmt.Println("Its New Date")
	}
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Its int")
	case float64:
		fmt.Println("Its float")
	case string:
		fmt.Println("Its string")
	default:
		fmt.Println("No idea whats its is")
	}
}
