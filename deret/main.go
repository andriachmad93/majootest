package main

import "fmt"

var (
	adder  = 0
	numb   = 0
	result = []int{}
)

func Deret(num1 int, num2 int, x int) []int {

	for i := 0; i < x; i++ {
		if i == 0 {
			numb = num1
		} else {
			adder = adder + (num2 - num1)

			numb = adder + num1
		}

		result = append(result, numb)
	}

	return result
}

func main() {
	fmt.Printf("%d", Deret(5, 8, 7))
}
