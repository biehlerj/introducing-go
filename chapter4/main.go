package main

import "fmt"

func main() {
	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//}
	//divideBy3()
	fizzBuzz()
}

func divideBy3() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	fizz := "Fizz"
	buzz := "Buzz"
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(fizz + buzz)
		} else if i%3 == 0 {
			fmt.Println(fizz)
		} else if i%5 == 0 {
			fmt.Println(buzz)
		} else {
			fmt.Println(i)
		}
	}
}
