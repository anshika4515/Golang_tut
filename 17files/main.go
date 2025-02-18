package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// creating and writing in file
func main() {
	fmt.Println("Files learn")
	content := "This needs to go int the file - LearnYourWays"

	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

// read the file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//chnge to string
	fmt.Println("Text data inside file is\n", string(databyte))
}
