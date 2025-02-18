package main

import "fmt"

func main() {
	fmt.Println("Lets learn abt loops")
	days := []string{"sunday", "tuesday", "wednesday", "friday"}
	fmt.Println(days)

	//for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//loop using range
	for i := range days {
		fmt.Println(days[i])
	}

	//type of for each
	for index, day := range days {
		fmt.Println("index is %v and value is %v\n", index, day)
	}

	//another loop method
	rougueValue := 1
	for rougueValue < 10 {
		fmt.Println("Value is", rougueValue)
		rougueValue++
	}
}

//we can also use continue and break here
