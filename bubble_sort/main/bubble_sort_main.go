package main

import (
	"fmt"
	"go_basics/bubble_sort"
	"go_basics/linkedlist"
)

func main() {
	fmt.Println("Bubble Sort Example")
	ll := linkedlist.New()
	ll.Add(1)
	ll.Add(10)
	ll.Add(4)
	ll.Add(2)
	ll.Print()
	l := bubble_sort.Sort(ll)
	l.Print()
}
