package problems

import "fmt"

func Prime() {
	var n int
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("It is not a prime number.")
		return
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println("It is not a prime number.")
			return
		}
	}
	fmt.Println("It is a prime number.")
}
