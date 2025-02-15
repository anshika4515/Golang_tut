package main

import "fmt"

const LoginToken = "fiohreiuh" //const public

func main() {
	var username string = "anshika"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.48445
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//no var keyword
	users := 349504
	fmt.Println(users)

	fmt.Println(LoginToken)
}
