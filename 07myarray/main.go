package main

import "fmt"
func main(){
	fmt.Println("Learn abt arrays")

	var fruitList [4]string  //array declaration
	fruitList[0] = "apple"
	fruitList[1] = "kiwi"
	fruitList[3] = "banana"
	//give space if any index data is not defined
	fmt.Println("fruits array" , fruitList) 

	var vegList = [5]string{"potato","peas","tomato"}
	fmt.Println(vegList)
}