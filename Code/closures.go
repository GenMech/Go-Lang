package main

import "fmt"

// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
func intSeq() func() int { 
    i := 0
    return func() int {
        i++
        return i
    }
}

func closures() {

	// Creating a closure by calling intSeq()
    nextInt := intSeq()

	fmt.Println("Result:", nextInt())

	// Invoking the closure and printing the result
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

	// Creating a new closure with independent state
    newInts := intSeq()

	// 
    fmt.Println(newInts())
}

// Here, intSeq() is a function call, and it returns another function. The returned function is assigned to the variable nextInt. So, nextInt is indeed a variable, but it's a variable that holds a function.

// In Go, functions are first-class citizens, meaning they can be assigned to variables, passed as arguments to other functions, and returned from functions. The concept of closures allows functions to capture and carry the surrounding state with them, which is what seems to be happening in your example.