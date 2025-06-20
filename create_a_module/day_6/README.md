<!-- Return a random greeting -->
# Return a random greeting

Step 1: Do what is done in day_5

Step 2: Update `greetings.go` in `greetings` directory

```go
package greetings

import (
  "errors"
  "fmt"
  "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
  // If no name was given, return an error with a message.
  if name == "" {
    return name, errors.New("empty name")
  }
  // Create a message using a random format.
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
  // A slice of message formats.
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }

  // Return a randomly selected message format by specifying
  // a random index for the slice of formats.
  return formats[rand.Intn(len(formats))]
}
```

Step 3: Update `hello.go` in `hello` directory

```go
package main

import (
  "fmt"
  "log"

  "example.com/greetings"
)

func main() {
  // Set properties of the predefined Logger, including
  // the log entry prefix and a flag to disable printing
  // the time, source file, and line number.
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  // Request a greeting message.
  message, err := greetings.Hello("Gladys")
  // If an error was returned, print it to the console and
  // exit the program.
  if err != nil {
    log.Fatal(err)
  }

  // If no error was returned, print the returned message
  // to the console.
  fmt.Println(message)
}
```

Run the code in `hello` directory

```bash
go run .
```

## new module

```go
import "math/rand" /* generate random number for selecting an item from the slice */
```

## slice

```go
/* A slice is like an array, except that its size changes dynamically as you add and remove items. */

variable := []{dataType}{
  // items
}

```

Go to `hello` directory and run the code

```bash
go run .
```
