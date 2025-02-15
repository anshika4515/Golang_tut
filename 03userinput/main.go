package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")

	//comma ok syntax or error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
}

//Type of taking input
