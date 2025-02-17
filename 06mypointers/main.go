package main

import "fmt"

func main() {
    fmt.Println("Learn about pointers")
    
    // Declare the pointer
    var ptr *int
    fmt.Println(ptr)

    // Creating the pointer referencing another variable
    myNumber := 23
    ptr = &myNumber // Using assignment here, not redeclaration
    fmt.Println(ptr)
    fmt.Println(*ptr) // Dereferencing the pointer to print the value

	*ptr = *ptr + 2
	fmt.Println(myNumber)   //added +2 in myNumber value
} 

//pointers ->operation are performed on actual values not on copy of values