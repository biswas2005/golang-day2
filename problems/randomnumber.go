package problems

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1

	var number int
	for {
		fmt.Println("Enter a number: ")
		fmt.Scan(&number)

		if number < secret {
			fmt.Println("Too low")
		} else if number > secret {
			fmt.Println("Too high")
		} else {
			fmt.Println("Wohoo! You got it.")
			break

		}

	}

}
