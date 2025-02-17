package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")
    anshika := User{"anshika","anshika@go.dev",true , 16}
	fmt.Println(anshika)
	fmt.Printf("anshika details are: %+v\n", anshika)
	fmt.Printf(anshika.Name)
	
}

//declaring the stuct
type User struct {
	Name string
	Email string
	Status bool
	Age int
}