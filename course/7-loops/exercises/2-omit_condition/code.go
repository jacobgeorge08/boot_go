package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	count := 0
	cost := 0.0
	for fee := 0.00; ; fee += 0.01 {
		if cost >= thresh {
			return count - 1
		}
		cost += 1 + fee
		count++
	}
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
	test(50.00)
}
