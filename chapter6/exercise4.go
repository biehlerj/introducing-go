package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	odd := makeOddGenerator()
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
}
