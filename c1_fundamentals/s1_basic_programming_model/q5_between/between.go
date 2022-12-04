package q5_between

import "fmt"

func between(x, y float64) {
	if x > 0 && x < 1 && y > 0 && y < 1 {
		fmt.Println(true)
	}

	fmt.Println(false)
}
