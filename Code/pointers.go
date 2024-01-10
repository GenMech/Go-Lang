package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func pointers() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}

// This function zeroval takes an integer parameter ival by value. It means that it works with a copy of the original value, and any changes made to ival inside the function do not affect the original variable outside the function.

// This function zeroptr takes an integer pointer iptr as a parameter. The *int indicates that it expects a memory address of an integer (a pointer to an integer). Inside the function body, *iptr = 0 dereferences the pointer, meaning it accesses the value stored at the memory address pointed to by iptr and sets that value to 0. This change will affect the original variable outside the function because you are modifying the value at the memory address it points to.