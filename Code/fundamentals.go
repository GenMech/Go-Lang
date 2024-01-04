package main

import "fmt"

func main() {
    fmt.Println("hello world")
    basic()
}

//The Go fmt package supports two closely-related functions for formatting a string to be displayed on the terminal. Print() accepts strings as arguments and concatenates them without any spacing. . Println() , on the other hand, adds a space between strings and appends a new line to the concatenated output string


// For loop in Go
func basic(){
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 4 ; j <= 10; j++{
		fmt.Println("Second loop number is:", j)
	}

	for{
		fmt.Println("This for loop without condition will run repeatedly until I break out of loop or return from enclosing fucntion")
		break
	}

}