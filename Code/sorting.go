package main

import (
	"fmt"
	"slices"
)

func sorting() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

//  Go’s slices package implements sorting for builtins and user-defined types.
//  Sorting functions are generic, and work for any ordered built-in type
