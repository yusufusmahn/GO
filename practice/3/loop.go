package main

import "fmt"

func main() {
	number := 1
	for number <= 100 {
		fmt.Println("Current number:", number)
	number *= 2
	}
}