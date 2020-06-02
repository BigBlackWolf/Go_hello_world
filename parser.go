package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://habr.com/ru/flows/develop/page%d/"
	var urls [50]string

	for i := 1; i <= 50; i++ {
		tmp := fmt.Sprintf(url, i)
		// urls = append(urls, tmp)
		urls[i-1] = tmp
		fmt.Println(tmp)
	}

	fmt.Println(urls)
	var result map[string]interface{}

	resp, _ := http.Get("https://api.ipify.org/?format=json")
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(a)
	json.Unmarshal(a, &result)
	fmt.Println(result)
}
