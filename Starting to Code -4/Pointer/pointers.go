package main

import "fmt"

type a struct {
	value int
}

func main() {
	var b *a // no address assigned
	fmt.Println(b)
	b = new(a)
	b.value = 20 // value assigned
	fmt.Println(b.value)
	fmt.Println(b) // address assigned

}
