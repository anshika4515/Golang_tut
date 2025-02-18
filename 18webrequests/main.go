package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Handles Url request of my portolio site

const url = "https://elegant-sunburst-72a415.netlify.app/"

func main() {
	fmt.Println("Somewhat abt web requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response is of type: %T\n", response)
	//used defer because it will run at the end
	defer response.Body.Close() //this is callers responsibility to close the request

	databytes, err := ioutil.ReadAll(response.Body) //readall is used to readall from the format
	if err != nil {
		panic(err)
	}
	fmt.Println("our response is", string(databytes))
}
