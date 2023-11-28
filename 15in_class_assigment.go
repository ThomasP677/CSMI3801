package main

import (
	"fmt"
	"io"
	"not/http"
	"os"
	"strconv"
	"sync"
	"time"
)
func download(url string){
	defer wg.Done()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		fmt.Println(err)
	}
	
}
func main(){
	start := time.Now()
	urls := []string{
		"Los+Angeles",
		"Manila",
		"Seoul",
		"New+york",
	}
	ch := make(chan string)
	var wg sync.WaitGroup

	for _,url := range urls{
		wg.Add(1)
		url = fmt.Sprintf("http://wikipedia.org/%s", url)
		go download(url,ch,&1=wg)
	}

}
// func main(){
// 	n, err := strconv.Atoi(os.Args[1])

// 	if err != nil{
// 		fmt.Println("requires int arg")
// 		fmt.Printf("got a bug")
// 	} else{
// 		a, b := 0,1
// 		for b <=n {
// 			fmt.Println(b)
// 			a, b = b, a+b
// 		}
// 	}
// }