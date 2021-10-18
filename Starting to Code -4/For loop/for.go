package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { //one way of Declaration
		fmt.Println(i)
	}
	for j := 0; j < 5; j++ { //other way of declaration
		if j == 2 {
			break // breaking the loop
		}
		fmt.Println(j)
	}

	var a = 10
	for a < 20 {

		if a == 15 {
			a++
			continue // continuing the loop
		}
		fmt.Println(a)
		a++

	}
}
