package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	url := "https://habr.com/ru/flows/develop/page%d/"
	var urls [50]string

	for i := 1; i <= 50; i++ {
		tmp := fmt.Sprintf(url, i)
		urls[i-1] = tmp
	}

	var wg sync.WaitGroup

	for counter, url := range urls {
		wg.Add(1)
		go request(url, wg, counter)
		counter++
	}

	// wg.Wait()
	time.Sleep(time.Second * 2)
	fmt.Println("Я все")
}

func request(url string, wg sync.WaitGroup, number int) {
	defer wg.Done()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var stringify string = string(body)
	go writeToFile(stringify, number)
	result := fmt.Sprintf("Done %s", url)
	fmt.Println(result)
}

func writeToFile(data string, number int) {
	filename := fmt.Sprintf("page%d.html", number)
	outFile, _ := os.Create(filename)
	defer outFile.Close()

	toWrite := []byte(data)
	outFile.Write(toWrite)
}

func read_from_file() {
	ioutil.ReadFile("site_list.txt")
}
