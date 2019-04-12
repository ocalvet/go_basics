package main

import (
	"fmt"

	"go_basics/linkedlist"
)

func main() {
	fmt.Println("Link list example")
	list := linkedlist.New()
	list.Add(4)
	list.Add(23)
	list.Print()
}
