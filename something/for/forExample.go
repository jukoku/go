package forexample

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}

func ForExample() {
	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(result)
}
