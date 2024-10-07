package main

import (
	"fmt"

	greetings "example.com/greetingslib"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.SayHello("GoLang")
	fmt.Println(message)

	// print Goodbye message
	fmt.Println(greetings.SayGoodbye("GoLang"))
}
