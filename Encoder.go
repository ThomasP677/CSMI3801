package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func fetchWiki(){
	var data []struct{
		State string `json:"state-province"`
		Country string `json:"country"`
		Domains []string `json:"country"`
		WebPages []string `json:"web_pages"`
		AlphaCode string `json:"alpha_two_code"`
		Name string `json:"name"`
	}
	url := fmt.Sprintf("http://universities.hipolabs.com/search?name=Loyola+Marymount+University")

	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(url)

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil{
		fmt.Printf("Error")
		return 
	}
	fmt.Printf("decoded data: %+v", data)
}

func main(){
	fetchWiki()
}