package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	fmt.Println("Hello")
}

//defer -> runs at last
//defer works on Last in First Out
