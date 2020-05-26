package main

import "fmt"

func main() {
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
