package main

import "fmt"

func greatestNumber(args ...int) int {
	greatest := 0

	for _, v := range args {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

func main() {
	fmt.Println(greatestNumber(5, 10, 19234, 111883, 20))
}
