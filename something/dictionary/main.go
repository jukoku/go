package main

import (
	"fmt"

	"github.com/jukoku/leargo/something/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
	err2 := dictionary.Delete(baseWord)
	if err2 != nil {
		fmt.Println(err2)
	}
}
