package problems

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rock = iota
	paper
	scissor
)

func Rps() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	computerChoice := rand.Intn(3)

	var userChoice int
	fmt.Println("<Enter 0 for rock><Enter 1 for paper><Enter 2 for scissor>")
	fmt.Println("Enter your choice: ")
	fmt.Scan(&userChoice)
	if (userChoice < 0) || (userChoice > 2) {
		fmt.Println("Invalid choices.")
		return
	}
	fmt.Println("Computer chose: ", computerChoice)

	if userChoice == computerChoice {
		fmt.Println("Its a draw!")
		//User winning probabalities
	} else if (userChoice == 0 && computerChoice == 2) || (userChoice == 1 && computerChoice == 0) || (userChoice == 2 && computerChoice == 1) {
		fmt.Println("You win.")
		//Computer wins
	} else {
		fmt.Println("You loose..")
	}
}
