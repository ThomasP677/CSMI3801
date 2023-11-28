package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)
func download(url string, ch chan <- string, wg * sync.WaitGroup){
	defer wg.Done()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil{
		fmt.Println(err)
	}
	bodyBytes, err :=io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
func main(){
	start := time.Now()
	urls := []string{
		"Los+Angeles",
		"Manila",
		"Seoul",
		"New+York",
	}
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, url := range urls {
		url = fmt.Sprintf("http://wikipedia.org/%s", url)
		wg.Add(1)
		go download(url, ch, &wg)
	}
	go func () {
		wg.Wait()
		close(ch)
	}()

	for result := range ch{
		fmt.Println(result)
	}
	fmt.Println("Time: ", time.Since(start))
}

	
