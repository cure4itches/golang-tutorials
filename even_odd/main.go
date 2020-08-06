package main

import "fmt"

func main() {
	nums := []int{}

	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	for i := range nums {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
