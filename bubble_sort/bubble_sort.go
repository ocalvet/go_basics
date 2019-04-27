package bubble_sort

import (
	"fmt"
	"go_basics/linkedlist"
)

type Iter func() (interface{}, bool)

func Sort(list *linkedlist.LinkedList) linkedlist.LinkedList {
	// m := linkedlist.New()
	iter := list.Iterator()
	for v, hasNext := iter(); hasNext; {
		if v != nil && hasNext {
			fmt.Println(v)
		}
		v, hasNext = iter()
	}
	return *linkedlist.New()
}
