package main

import (
	"fmt"
)

func main() {
	// map[key type]value type
	nico := map[string]string{"name": "nico", "age": "12"}
	for _, value := range nico {
		fmt.Println(value)
	}
}
