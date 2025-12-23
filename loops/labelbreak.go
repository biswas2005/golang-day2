package loops

import "fmt"

func Labelbreak() {

outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j <= 1 {
				break outer
			}
			fmt.Println(i, j)
		}
	}
}
