package main

import "fmt"

func main() {
	fmt.Println("Lets learn abt functions")
	greeter() //basic function execution

	result := adder(3, 5)
	fmt.Println("result is", result)
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func greeter() {
	fmt.Println("Namste from golang")
}
