package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	//for time
	fmt.Println(presentTime)
	//format for time n date
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //give the correct date even if you enter wrong here

	//if we want to create our own time
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
