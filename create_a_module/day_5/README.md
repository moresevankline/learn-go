<!-- Return and handle an error -->
# Return and handle an error

Step 1: Do what is done in Day 4

Step 2: Update `greetings.go` in `greetings` directory

```go
package greetings

import (
  "errors"
  "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
  // If no name was given, return an error with a message.
  if name == "" {
    return "", errors.New("empty name")
  }

  // If a name was received, return a value that embeds the name
  // in a greeting message.
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}
```

Step 3: Update `hello.go` in `hello` directory

Go to `hello` directory and run the code

```bash
go run .
```

## if statement in Go

```go
if {condition} {
  // code execution
}
```

## modules and their purpose

```go
import "errors" /* for error handling */
import "log" /* for log messages */
```
