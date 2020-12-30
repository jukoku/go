package main

import (
	"fmt"

	"github.com/jukoku/leargo/something/structs/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
