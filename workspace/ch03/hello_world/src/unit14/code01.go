package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("CPU:", runtime.NumCPU())
}
