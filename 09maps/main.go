package main

import "fmt"
func main(){
	fmt.Println("Learn abt maps")
	//declare map using make

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	//fetch individual value from map
	fmt.Println(languages["JS"])

	//delete from map
	delete(languages,"RB")
	fmt.Println(languages)

	//loops in golang
	for key, value := range languages {
		fmt.Println("For key %v, value is %v\n", key, value)
	}
}