package main

import (
	"fmt"
	"maps"
)

func main() {
    fmt.Println("hello world")
    basic() // for loop basics
	array() // basic of arrays
	hashmap() // hashmaps basics
    ranges() // for ranges
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

// Hashmap in GO
func hashmap() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    clear(m)
    fmt.Println("map:", m)

    _, prs := m["k2"] // Here,we are using the blank identifier (_) to ignore the first value returned by the expression m["k2"]. This is commonly used when you are interested in checking the presence of a key in a map but don't need the actual value associated with that key.
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}

//The _ (underscore) is known as the blank identifier in Go. We can use the blank identifier to declare and use the unused variables. The unused variables are the variables that are only defined once in the program and not used. 

// Ranges in Go
func ranges(){
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums { // Here we dont want index so we ignored it
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}