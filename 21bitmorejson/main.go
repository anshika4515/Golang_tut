package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	platform string
	Password string
	Tags     []string //slice
}

func main() {
	fmt.Println("More abt JSON")
	//EncodeJson()
	DecodeJson()
}

// convert the data to the json data

func EncodeJson() {
	lcocourses := []course{
		{"Anshika", 234, "Learn.com", "abchd67", []string{"reactjs", "web-dev"}},
		{"Verma", 738, "eco.in", "dieoij", []string{"javascript", "ejkdm"}},
	}

	//change to Json
	// content, err := json.Marshal(lcocourses)
	content, err := json.MarshalIndent(lcocourses, "", "\t") //gives in proper format
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs Bootcamp",
			"Price": 299,
			"website": "learn.in",
			"tags": ["web-dev", "js"]
		}
	 `)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb) //check if my json is valid or not

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //used to avoid the json
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json not valid")
	}
}
