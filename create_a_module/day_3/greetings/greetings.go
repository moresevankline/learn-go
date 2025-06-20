package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

/*
	In Go, the := operator is a shortcut for declaring,
	and initializing a variable in one line

	Long way:
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
*/

/*
	Use the fmt package's Sprintf function to create a greeting message.
	The first argument is a format string,
	and Sprintf substitutes the name parameter's value for the %v format verb.
*/
