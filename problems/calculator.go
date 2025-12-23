package problems

import (
	"fmt"
	"math"
)

func Calculator() {
	var num1, num2 float64
	var operator string
	fmt.Println("Number1 <Operator> Number2")
	fmt.Scan(&num1, &operator, &num2)

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	case "%":
		fmt.Println(math.Mod(num1, num2))
	}
}
