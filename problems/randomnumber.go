package problems

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	secret := rand.Intn(100) + 1

	maxattempts := 5
	fmt.Println("Guess a number between 0 and 100")
	fmt.Println("You have", maxattempts, "to guess")
	var number int
	attempts := 0
	for attempts < maxattempts {
		fmt.Println("Enter a number: ")
		fmt.Scan(&number)
		attempts++
		if number < secret {
			fmt.Println("Too low")
		} else if number > secret {
			fmt.Println("Too high")
		} else {
			fmt.Println("Wohoo! You are correct")
			return
		}
	}
	fmt.Println("You loose!")
	fmt.Println("The correct number was: ", secret)
}
