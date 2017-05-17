package main

import "fmt"
import "unicode/utf8"

func main(){
    	var s1 = "hello"
	var s2 = "김태완"
	var s3 = "김 태완"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(len(s3))

	fmt.Println("---------------")

	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Println(utf8.RuneCountInString(s2))
	fmt.Println(utf8.RuneCountInString(s3))
}


