package main

import "fmt"

func main() {
	shopping := map[string]int{
		"computer": 100,
		"laptop":   200,
		"mobile":   300,
	}
	for k, v := range shopping {
		fmt.Println(k, v)
	}
	for k := range shopping {
		fmt.Println(k) // printing only the key
	}
	for _, v := range shopping {
		fmt.Println(v) //printing only the values
	}

}
