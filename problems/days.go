//Guess the day in a week based on number by user

package problems

import "fmt"

func Days() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	switch number {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")

	}
}
