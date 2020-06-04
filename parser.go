package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var dataFromFile []string = read_from_file("site_list.txt")
	fmt.Println(len(dataFromFile))
	var urls [50]string = generate_urls()
	parse(dataFromFile)
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

func read_from_file(file string) []string {
	var lines []string
	data, _ := ioutil.ReadFile(file)
	readable := string(data)
	fmt.Println(readable)

	f, _ := os.Open(file)
	defer f.Close()

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	return lines
}

func generate_urls() [50]string {
	url := "https://habr.com/ru/flows/develop/page%d/"
	var urls [50]string

	for i := 1; i <= 50; i++ {
		tmp := fmt.Sprintf(url, i)
		urls[i-1] = tmp
	}
	return urls
}

func parse(urls []string) {
	var wg sync.WaitGroup

	for counter, url := range urls {
		wg.Add(1)
		go request(url, wg, counter)
	}

	// wg.Wait()
	time.Sleep(time.Second * 3)
	fmt.Println("Я все")
}
