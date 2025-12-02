package basics

import "fmt"

// letsSing:="Yo Yo honey singer"
var letsSing = "Yo Yo honey singer"

func main3() {
	// Ha ha ha . I am comment. I won't execute yet I exists

	var age int
	var name string = "Aaradhya" //Lol thats Amitabh bachan grand daughter name
	var surname = "Gopal"

	fmt.Println(age, name, surname)
	count := 5
	course := "Go Bootcamp" //Hi hi ha ha ha; another bootcamp which you probably start but won't complete
	fmt.Println(count, course)
	printName()
	// benFirstName//What is this??
	fmt.Println(letsSing)

}

func printName() {
	benFirstName := "Ben"
	benLastName := "10"
	fmt.Println(benFirstName, " ", benLastName)
	fmt.Println(letsSing)
}
