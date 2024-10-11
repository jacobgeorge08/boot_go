package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if e.isSubscribed == true {
		return float64(len(e.body)) * 0.02
	} else {
		return float64(len(e.body)) * 0.05
	}
}

func (e email) print() {
	if e.isSubscribed == true {
		fmt.Printf("'%v' | Subscribed\n", e.body)
	} else {
		fmt.Printf("'%v' | Not Subscribed\n", e.body)
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}
