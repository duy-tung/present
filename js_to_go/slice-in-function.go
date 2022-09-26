package main

import "fmt"

func once(x [3]int) {
	for i := range x {
		x[i] *= 2
	}
}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

func main() {
	data := [3]int{8, 9, 0}

	once(data)
	fmt.Println(data)

	twice(data[0:])
	fmt.Println(data)
}
