package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	anshika := User{"anshika", "anshika@go.dev", true, 16}
	fmt.Println(anshika)
	fmt.Printf("anshika details are: %+v\n", anshika)
	fmt.Printf(anshika.Name)
	anshika.GetStatus() //declaring method
	anshika.NewMail()

}

// declaring the stuct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// abt methods
func (u User) GetStatus() {
	fmt.Println("is user active", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println(u.Email)
}
