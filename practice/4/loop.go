package main

import "fmt"
func main() {
	for{
		fmt.Print("enter a positive number or -1 to exit: ")
		var num int
		fmt.Scan(&num)
		if num < 0 {
			fmt.Println("Exiting the loop.")
			break
		}
		square := num * num
		// fmt.Println("Square of", num, "is", square)
		fmt.Printf("Square of %d is %d\n", num, square)

	}
}