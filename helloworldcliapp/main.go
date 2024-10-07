package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(quote.Go())

	whatAmI := "a newbie GoLang Programmer"

	fmt.Println("I am", whatAmI)
}
