package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza App")
	fmt.Println("Please rate our pizza between 1 and 5")
	//taking input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') //read string until it hits space

	fmt.Println("Thanks for rating", input)

	//convert string to int to add 1
	//used like try n catch
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating", numRating+1)
	}
}
