package main

import ("fmt"
         "sort" )
func main(){
	fmt.Println("Slices")
	// []string  -> this is type of slices
	var fruitList = []string{"apple","kiwi","banana"}  //declre slice
	fruitList = append(fruitList,"orange","cherry")   //add in slice
	fmt.Println("fruit list is" , fruitList)

	fruitList = append(fruitList[1:3])  //used to show the range within the slice
	fmt.Println(fruitList)

	//use of make 
	//make(data type , size)
	highscores := make([]int ,4)  //default value of memory
	highscores[0] =273
	highscores[1] = 278
	highscores[2] = 837
	highscores[3] = 829

	highscores = append(highscores, 235 , 899)  //add the data dispite of the memory defined
	sort.Ints(highscores)   //sort the integer values
	fmt.Println("highscore is" , highscores)

	//how to remove value from slice based on index?
	var courses = []string{"reactjs","aws","go","javascript"}
	var index int =2
	courses = append(courses[:index],courses[index+1:]...)   //remove go from array
	fmt.Println(courses)
}