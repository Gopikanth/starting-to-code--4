package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("Middle") // executed only at last
	fmt.Println("End")
}
