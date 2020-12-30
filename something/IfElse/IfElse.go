package ifelse

import (
	"fmt"
)

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge <= 19 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
