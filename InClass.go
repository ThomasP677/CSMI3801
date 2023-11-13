package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	n,err := strconv.Atio(os.Args[1])

	if err != nil{
		fmt.Println("requires int arg")
		fmt.Printf("got bug: %s \n", err)
	} else{
		a,b := 0,1
		for b <= n{
			fmt.Println(b)
			a, b = b, a+b
		}
	}
}