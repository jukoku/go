package main

import (
	"fmt"

	"github.com/go/something/structs/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
