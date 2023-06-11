package main

import (
	"fmt"
)

func isPos(x int) bool {
	if x >= 0 {
		return true
	} else {
		return false
	}

}

func isEven (x int) bool {
	if y := x; y % 2 == 0 {
		return true
	}
	return false
}

func main() {
	sum := 0
	for sum < 1000 {
		sum += 1
	}
	fmt.Println(sum)
	fmt.Println(isPos(-2))
	fmt.Println(isEven(4))
	fmt.Println(isEven(5))
}
