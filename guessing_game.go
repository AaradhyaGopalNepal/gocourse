package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1
	fmt.Println("Welcome to gussing game")
	fmt.Println("I have number from 1 to 100, try and guess the number")
	var guess int
	for {
		fmt.Println("Enter a number")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("Congratulation you guessed correct number")
			break
		} else if guess < target {
			fmt.Println("Try guess higher number")
		} else {
			fmt.Println("Try guess lower number")
		}
	}
}
