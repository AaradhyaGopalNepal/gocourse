package basics

import "fmt"

const PI = 3.14
const GRAVITY float32 = 9.81

func main5() {

	const (
		MONYAYYY    = 1
		TUESYAYYY   = 2
		WEDNESYAYYY = 3
	)
	const THURSYAYYY = WEDNESYAYYY + MONYAYYY
	fmt.Println("Hello PI value is :", PI, "and gravity value is:", GRAVITY)
	fmt.Println("Monday", MONYAYYY, "Tuesday", TUESYAYYY, "Wednesday", WEDNESYAYYY, "Thursday", THURSYAYYY)
}
