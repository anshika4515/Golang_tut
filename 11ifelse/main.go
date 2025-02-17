package main

import "fmt"

func main(){
	fmt.Println("Learn abt if n else")

	//abt if else condition
    loginCount := 23
	var result string

	if loginCount < 10{
		result = "small"
	}else if loginCount >10 {
		result = "large"
	}else{
		result ="Exactly 10"
	}

	fmt.Println(result)

} 