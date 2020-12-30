package main

import (
	"fmt"
)

func main() {
	a := 2
	// & -> adress
	b := &a
	// * saved value of adress
	*b = 20
	fmt.Println(*b)
}
