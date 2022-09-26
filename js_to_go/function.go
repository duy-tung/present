package main

import "fmt"

func doThings(args ...int) int {
	total := 0
	for _, num := range args {
		total += num
	}
	return total
}

func main() {
	fmt.Println(doThings(1, 2, 3))
}
