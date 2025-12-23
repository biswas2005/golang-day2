package loops

import "fmt"

func Labelcontinue() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue outer
			}
			fmt.Println(i, j)
		}
	}
}
