<!-- Return greetings for multiple people -->
# Return greetings for multiple people

Step 1: Do what is done in Day 6

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
    return "", errors.New("empty name")
  }

  // If a name was received, return a value that embeds the name
  // in a greeting message.
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
  // A map to associate names with messages.
  messages := make(map[string]string)
  // Loop through the received slice of names, calling
  // the Hello function to get a message for each name.
  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }
    // In the map, associate the retrieved message with
    // the name.
    messages[name] = message
  }
  return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
  // A slice of message formats.
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v",
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

  // A slice of names.
  names := []string{"Gladys", "Samantha", "Darrin"}

  // Request greeting messages for the names.
  messages, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  // If no error was returned, print the returned map of
  // messages to the console.
  fmt.Println(messages)
}

```

Go to `hello` directory and run the code

```bash
go run .
```
