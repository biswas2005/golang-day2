package problems

import "fmt"

func Fibonacci() {

	var n int
	fmt.Println("Enter limit you want: ")
	fmt.Scan(&n)

	a := 0
	b := 1

	fmt.Println("Fibonacci series is: ")
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
