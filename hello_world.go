package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func multiply(a int, c int) (result int, comment string) {
	return a*c, "hello"
}

func main() {
	sum, _ := multiply(3, 4)
	fmt.Println(sum)
	array_testing()
	pointer, pointer2 := memory_testing()
	fmt.Println(pointer, *pointer2)
	lambda_function()

	structure_testing()
	args_testing(1, "asd")
	concurency_testing()
	webProgramming_testing()
}

func array_testing(){
	var first [5]string // just define -> array
	second := []int{1, 2, 3} // initialize -> slice (not fixed size)

	second = append(second, 200)
	second = append(second, []int{400, 500}...)
	second[len(second) - 1] = 100
	
	fmt.Println(first)
	fmt.Println(second)
}

func memory_testing() (p, q *int) {
	p = new(int) // allocate memory
	s := make([]int, 10) // allocate memory -> array
	s[3] = -1
	r := 1
	return &r, &s[3]
}

func dictionary_testing() map[string]int {
	dict := map[string]int{"three": 3, "one": 1}
	dict["second"] = 2 
	return dict
}

func lambda_function() {
	fmt.Println("Add two numbers: (2 + 7) * 2 = ", 
		func(a, b int) int { 
			return (a + b) * 2
		}(2, 7))
}

type Node struct {
	value int
	next *Node
}

type LinkedList interface {
	nextNode()
}

func (self Node) nextNode(next *Node) {
	self.next = next
	fmt.Println(self.next.value)
}

func structure_testing() {
	a := Node{1, nil}
	b := Node{2, nil}
	a.nextNode(&b)
	fmt.Println(a.next)
}

func args_testing(args ...interface{}){
	for _, param := range args {
		fmt.Println("param: ", param)
	}
}

func inc(i int, c chan int) {
	c <- i + 1
}

func concurency_testing() {
	c := make(chan int)
	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)

	fmt.Println(<-c, <-c, <-c)
}

type pair struct {
	x, y int
}

func webProgramming_testing() {
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err)
	}()

	requestServer()
}

func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world response via GO!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\n Webserver said: %s", string(body))
}

func test() {
	var a int = 1
	var b string = "1q3ew"
	c := "test"
	fmt.Println(a, b, c) 

	i := 1
	for i <= 3 {
		fmt.Println("i: ", i)
		i = i + 1
		if i == 3 {
			break
		}
	}
	
	for j := -2; j <= 0; j++{
		fmt.Println("j: ", j)
	}

	if num := 10; c == "tes" {
		a = 5	
	} else if num < 0 {
		c = "s"
	} else {
		c = "nbfd"
	}

	switch a {
		case 0:
			fmt.Println("Zero")
		case 2, 1:
			fmt.Println("One or two")
		default:
			fmt.Println("Non of them")
	}
	
	var array[5] int
	array[2] = 1
	array2 := [2]string{"cyber", "punk"}
	fmt.Println("array 2:", array2)

	array_copy := array
	array_copy[0] = 123
	fmt.Println(array[0] == array_copy[0])
}
