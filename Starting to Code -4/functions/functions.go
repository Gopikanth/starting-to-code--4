package main

import "fmt"

func main() {
	msg := "hai"
	second(msg)
	fmt.Println(msg) //this prints raju
}
func second(msg string) {
	msg = "raju"
	fmt.Println(msg) //this prints hai
}
