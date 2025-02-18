package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Learn abt get request in go lang")
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response", response)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"
	//fake json payload
	requestBody := strings.NewReader(`
	     {
	         "coursename":"Lets learn go with golang",
			 "price":0,
			 "platform":"learnCodeOnline.in"
	     }
	`)
	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"
	//to extract form data
	data := url.Values{}
	data.Add("firstname", "anshika")
	data.Add("lastname", "verma")
	data.Add("email", "anshika@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	//to read the response
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
