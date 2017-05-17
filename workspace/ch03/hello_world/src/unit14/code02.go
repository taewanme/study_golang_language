package main

import (
	. "fmt"
	"runtime"
)

func main(){
	Println("CPU:", runtime.NumCPU())
}
