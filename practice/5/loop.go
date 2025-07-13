package main
import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	for i, num:= range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}