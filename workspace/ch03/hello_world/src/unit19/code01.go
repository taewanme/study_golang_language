package main

import "fmt"

func main(){
	var a = 3
	switch {
	case a >= 5 && a< 10:
		fmt.Println("5~9")
	case a >=0 && a<5:
		fmt.Println("0~4")
	}


}
