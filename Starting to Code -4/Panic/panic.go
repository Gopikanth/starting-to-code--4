package main

import "fmt"

func main() {
	fmt.Println("start")
	panic("Panic happended") //shows panic happened
	fmt.Println("End")
}
