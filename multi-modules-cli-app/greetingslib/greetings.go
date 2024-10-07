package greetings

import "fmt"

// SayHello return a Hello greeting for the named person.
func SayHello(name string) string {
	// Return a formatted greeting message that says "Hello..." with the given name
	// var message string                       // variable declaration
	// message = fmt.Sprintf("Hello %v!", name) // variable initialization

	message := fmt.Sprintf("Hello %v!", name) // shorthand variable declaration and init
	return message
}

// Returns a Goodbye greeting message for the named person.
func SayGoodbye(name string) string {
	return fmt.Sprintf("Goodbye %v", name)
}
