package loops

import "fmt"

func Simplefor() {
	for i, j := 0, 8; i < j; i, j = i+1, j-1 {
		{
			fmt.Println(i, j)
		}
	}
}
