package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for k, v := range a {
		fmt.Println(k, v) //k is the index & and v is the value
	}
}
