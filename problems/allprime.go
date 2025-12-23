package problems

import "fmt"

func Allprime() {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("Invalid range")
		return
	}
	fmt.Println("Prime numbers are: ")
	for n := num1; n <= num2; n++ {
		if isprime(n) {

			fmt.Println(n)

		}

	}
}
func isprime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
